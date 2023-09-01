package components

import (
	"math"
	"time"

	policylangv1 "github.com/fluxninja/aperture/v2/api/gen/proto/go/aperture/policy/language/v1"
	"github.com/fluxninja/aperture/v2/pkg/config"
	"github.com/fluxninja/aperture/v2/pkg/notifiers"
	"github.com/fluxninja/aperture/v2/pkg/policies/controlplane/iface"
	"github.com/fluxninja/aperture/v2/pkg/policies/controlplane/runtime"
	"go.uber.org/fx"
)

// PIDController .
type PIDController struct {
	lastSignal     runtime.Reading
	lastOutput     runtime.Reading
	parameters     *policylangv1.PIDController_Parameters
	integral       float64
	invalidCount   int
	ticksPerSample int
}

// Make sure PID complies with Component interface.
var _ runtime.Component = (*PIDController)(nil)

// Name implements runtime.Component.
func (*PIDController) Name() string { return "PID" }

// Type implements runtime.Component.
func (*PIDController) Type() runtime.ComponentType { return runtime.ComponentTypeSignalProcessor }

// ShortDescription implements runtime.Component.
func (*PIDController) ShortDescription() string { return "PID Controller" }

// IsActuator implements runtime.Component.
func (*PIDController) IsActuator() bool { return false }

// NewPIDControllerAndOptions creates a PID component and its fx options.
func NewPIDControllerAndOptions(pidProto *policylangv1.PIDController, _ runtime.ComponentID, policyReadAPI iface.Policy) (runtime.Component, fx.Option, error) {
	pid := &PIDController{
		parameters: pidProto.Parameters,
		lastSignal: runtime.InvalidReading(),
		lastOutput: runtime.InvalidReading(),
	}
	// period of tick
	evaluationPeriod := policyReadAPI.GetEvaluationInterval()
	var samplePeriod time.Duration
	if pidProto.Parameters.SamplePeriod == nil {
		samplePeriod = evaluationPeriod
	} else {
		samplePeriod = pidProto.Parameters.SamplePeriod.AsDuration()
	}
	//
	pid.ticksPerSample = int(math.Ceil(float64(samplePeriod) / float64(evaluationPeriod)))
	return pid, fx.Options(), nil
}

// Execute implements runtime.Component.Execute.
func (pid *PIDController) Execute(inPortReadings runtime.PortToReading, tickInfo runtime.TickInfo) (runtime.PortToReading, error) {
	if tickInfo.Tick()%pid.ticksPerSample != 0 {
		return runtime.PortToReading{
			"output": []runtime.Reading{pid.lastOutput},
		}, nil
	}

	signalVal := inPortReadings.ReadSingleReadingPort("signal")
	setpointVal := inPortReadings.ReadSingleReadingPort("setpoint")

	if !signalVal.Valid() || !setpointVal.Valid() {
		pid.invalidCount++
		if pid.invalidCount > int(pid.parameters.ResetAfterInvalidSamples) {
			pid.integral = 0
			pid.invalidCount = 0
			pid.lastOutput = runtime.InvalidReading()
		}
		return runtime.PortToReading{
			"output": []runtime.Reading{pid.lastOutput},
		}, nil
	}

	pid.invalidCount = 0
	signal := signalVal.Value()
	setpoint := setpointVal.Value()
	error := setpoint - signal
	derivative := math.NaN()

	// Calculate the proportional term
	proportional := error * pid.parameters.Kp

	// Calculate the integral term
	pid.integral += error * pid.parameters.Ki

	// Calculate the derivative term
	if pid.lastSignal.Valid() {
		derivative = (signal - pid.lastSignal.Value()) * pid.parameters.Kd
	}

	// Need a valid derivative to calculate the output
	if pid.parameters.Kd != 0 && math.IsNaN(derivative) {
		return runtime.PortToReading{
			"output": []runtime.Reading{runtime.InvalidReading()},
		}, nil
	}

	output := proportional + pid.integral + derivative
	outputPreConstraints := output

	maxVal := inPortReadings.ReadSingleReadingPort("max")
	if maxVal.Valid() {
		output = math.Min(output, maxVal.Value())
	}

	minVal := inPortReadings.ReadSingleReadingPort("min")
	if minVal.Valid() {
		output = math.Max(output, minVal.Value())
	}

	// To prevent the integral term from accumulating errors due to constraints on the output value
	if pid.parameters.Ki != 0 {
		correction := output - outputPreConstraints
		if correction != 0 {
			pid.integral += correction / pid.parameters.Ki
		}
	}

	outputVal := runtime.NewReading(output)
	pid.lastOutput = outputVal
	pid.lastSignal = signalVal

	return runtime.PortToReading{
		"output": []runtime.Reading{outputVal},
	}, nil
}

// DynamicConfigUpdate is a no-op for PID.
func (pid *PIDController) DynamicConfigUpdate(event notifiers.Event, unmarshaller config.Unmarshaller) {
}
