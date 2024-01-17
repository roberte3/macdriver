// Code generated by DarwinKit. DO NOT EDIT.

package fileprovider

import (
	"github.com/progrium/macdriver/objc"
)

// An operation that syncs the modification of the source item to the target location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingmodification?language=objc
type PFileProviderTestingModification interface {
	// optional
	ChangedFields() FileProviderItemFields
	HasChangedFields() bool

	// optional
	TargetItemIdentifier() FileProviderItemIdentifier
	HasTargetItemIdentifier() bool

	// optional
	SourceItem() objc.Object
	HasSourceItem() bool

	// optional
	TargetSide() FileProviderTestingOperationSide
	HasTargetSide() bool

	// optional
	DomainVersion() FileProviderDomainVersion
	HasDomainVersion() bool

	// optional
	TargetItemBaseVersion() FileProviderItemVersion
	HasTargetItemBaseVersion() bool
}

// ensure impl type implements protocol interface
var _ PFileProviderTestingModification = (*FileProviderTestingModificationObject)(nil)

// A concrete type for the [PFileProviderTestingModification] protocol.
type FileProviderTestingModificationObject struct {
	objc.Object
}

func (f_ FileProviderTestingModificationObject) HasChangedFields() bool {
	return f_.RespondsToSelector(objc.Sel("changedFields"))
}

// A list of the fields that changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingmodification/3736241-changedfields?language=objc
func (f_ FileProviderTestingModificationObject) ChangedFields() FileProviderItemFields {
	rv := objc.Call[FileProviderItemFields](f_, objc.Sel("changedFields"))
	return rv
}

func (f_ FileProviderTestingModificationObject) HasTargetItemIdentifier() bool {
	return f_.RespondsToSelector(objc.Sel("targetItemIdentifier"))
}

// The unique identifier for the target item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingmodification/3736245-targetitemidentifier?language=objc
func (f_ FileProviderTestingModificationObject) TargetItemIdentifier() FileProviderItemIdentifier {
	rv := objc.Call[FileProviderItemIdentifier](f_, objc.Sel("targetItemIdentifier"))
	return rv
}

func (f_ FileProviderTestingModificationObject) HasSourceItem() bool {
	return f_.RespondsToSelector(objc.Sel("sourceItem"))
}

// A description of the source item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingmodification/3736243-sourceitem?language=objc
func (f_ FileProviderTestingModificationObject) SourceItem() objc.Object {
	rv := objc.Call[objc.Object](f_, objc.Sel("sourceItem"))
	return rv
}

func (f_ FileProviderTestingModificationObject) HasTargetSide() bool {
	return f_.RespondsToSelector(objc.Sel("targetSide"))
}

// The target location for the modification operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingmodification/3736246-targetside?language=objc
func (f_ FileProviderTestingModificationObject) TargetSide() FileProviderTestingOperationSide {
	rv := objc.Call[FileProviderTestingOperationSide](f_, objc.Sel("targetSide"))
	return rv
}

func (f_ FileProviderTestingModificationObject) HasDomainVersion() bool {
	return f_.RespondsToSelector(objc.Sel("domainVersion"))
}

// The domain’s version when the change occurred. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingmodification/3736242-domainversion?language=objc
func (f_ FileProviderTestingModificationObject) DomainVersion() FileProviderDomainVersion {
	rv := objc.Call[FileProviderDomainVersion](f_, objc.Sel("domainVersion"))
	return rv
}

func (f_ FileProviderTestingModificationObject) HasTargetItemBaseVersion() bool {
	return f_.RespondsToSelector(objc.Sel("targetItemBaseVersion"))
}

// The version of the changed item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovidertestingmodification/3736244-targetitembaseversion?language=objc
func (f_ FileProviderTestingModificationObject) TargetItemBaseVersion() FileProviderItemVersion {
	rv := objc.Call[FileProviderItemVersion](f_, objc.Sel("targetItemBaseVersion"))
	return rv
}