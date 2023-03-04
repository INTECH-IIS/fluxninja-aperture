// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package entitycachev1

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using EntityCache within kubernetes types, where deepcopy-gen is used.
func (in *EntityCache) DeepCopyInto(out *EntityCache) {
	p := proto.Clone(in).(*EntityCache)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EntityCache. Required by controller-gen.
func (in *EntityCache) DeepCopy() *EntityCache {
	if in == nil {
		return nil
	}
	out := new(EntityCache)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new EntityCache. Required by controller-gen.
func (in *EntityCache) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using EntityCache_Entities within kubernetes types, where deepcopy-gen is used.
func (in *EntityCache_Entities) DeepCopyInto(out *EntityCache_Entities) {
	p := proto.Clone(in).(*EntityCache_Entities)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EntityCache_Entities. Required by controller-gen.
func (in *EntityCache_Entities) DeepCopy() *EntityCache_Entities {
	if in == nil {
		return nil
	}
	out := new(EntityCache_Entities)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new EntityCache_Entities. Required by controller-gen.
func (in *EntityCache_Entities) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Entity within kubernetes types, where deepcopy-gen is used.
func (in *Entity) DeepCopyInto(out *Entity) {
	p := proto.Clone(in).(*Entity)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Entity. Required by controller-gen.
func (in *Entity) DeepCopy() *Entity {
	if in == nil {
		return nil
	}
	out := new(Entity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Entity. Required by controller-gen.
func (in *Entity) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}