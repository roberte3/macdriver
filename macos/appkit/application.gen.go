// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Application] class.
var ApplicationClass = _ApplicationClass{objc.GetClass("NSApplication")}

type _ApplicationClass struct {
	objc.Class
}

// An interface definition for the [Application] class.
type IApplication interface {
	IResponder
	SendActionToFrom(action objc.Selector, target objc.IObject, sender objc.IObject) bool
	EndModalSession(session ModalSession)
	RequestUserAttention(requestType RequestUserAttentionType) int
	CancelUserAttentionRequest(request int)
	SendEvent(event IEvent)
	UpdateWindowsItem(win IWindow)
	ArrangeInFront(sender objc.IObject)
	RemoveWindowsItem(win IWindow)
	Hide(sender objc.IObject)
	RegisterServicesMenuSendTypesReturnTypes(sendTypes []PasteboardType, returnTypes []PasteboardType)
	TargetForActionToFrom(action objc.Selector, target objc.IObject, sender objc.IObject) objc.Object
	SearchStringInUserInterfaceItemStringSearchRangeFoundRange(searchString string, stringToSearch string, searchRange foundation.Range, foundRange *foundation.Range) bool
	StopModalWithCode(returnCode ModalResponse)
	UnregisterForRemoteNotifications()
	WindowWithWindowNumber(windowNum int) Window
	RunModalSession(session ModalSession) ModalResponse
	StopModal()
	AbortModal()
	ExtendStateRestoration()
	DisableRelaunchOnLogin()
	EnumerateWindowsWithOptionsUsingBlock(options WindowListOptions, block func(window Window, stop *bool))
	OrderFrontStandardAboutPanel(sender objc.IObject)
	CompleteStateRestoration()
	OrderFrontCharacterPalette(sender objc.IObject)
	DiscardEventsMatchingMaskBeforeEvent(mask EventMask, lastEvent IEvent)
	SetWindowsNeedUpdate(needUpdate bool)
	ActivationPolicy() ApplicationActivationPolicy
	ToggleTouchBarCustomizationPalette(sender objc.IObject) objc.Object
	FinishLaunching()
	PreventWindowOrdering()
	EnableRelaunchOnLogin()
	UnhideWithoutActivation()
	UnregisterUserInterfaceItemSearchHandler(handler PUserInterfaceItemSearching)
	UnregisterUserInterfaceItemSearchHandlerObject(handlerObject objc.IObject)
	RunPageLayout(sender objc.IObject)
	ReportException(exception foundation.IException)
	OrderFrontColorPanel(sender objc.IObject)
	Deactivate()
	AddWindowsItemTitleFilename(win IWindow, string_ string, isFilename bool)
	UpdateWindows()
	RegisterForRemoteNotifications()
	PostEventAtStart(event IEvent, flag bool)
	UnhideAllApplications(sender objc.IObject)
	ActivateContextHelpMode(sender objc.IObject)
	SetActivationPolicy(activationPolicy ApplicationActivationPolicy) bool
	MiniaturizeAll(sender objc.IObject)
	NextEventMatchingMaskUntilDateInModeDequeue(mask EventMask, expiration foundation.IDate, mode foundation.RunLoopMode, deqFlag bool) Event
	HideOtherApplications(sender objc.IObject)
	ReplyToOpenOrPrint(reply ApplicationDelegateReply)
	ShowHelp(sender objc.IObject)
	ReplyToApplicationShouldTerminate(shouldTerminate bool)
	Run()
	Stop(sender objc.IObject)
	Unhide(sender objc.IObject)
	RegisterUserInterfaceItemSearchHandler(handler PUserInterfaceItemSearching)
	RegisterUserInterfaceItemSearchHandlerObject(handlerObject objc.IObject)
	ChangeWindowsItemTitleFilename(win IWindow, string_ string, isFilename bool)
	Terminate(sender objc.IObject)
	RestoreWindowWithIdentifierStateCompletionHandler(identifier UserInterfaceItemIdentifier, state foundation.ICoder, completionHandler func(arg0 Window, arg1 foundation.Error)) bool
	OrderFrontStandardAboutPanelWithOptions(optionsDictionary map[AboutPanelOptionKey]objc.IObject)
	IsRegisteredForRemoteNotifications() bool
	ServicesMenu() Menu
	SetServicesMenu(value IMenu)
	IsRunning() bool
	IsHidden() bool
	KeyWindow() Window
	PresentationOptions() ApplicationPresentationOptions
	SetPresentationOptions(value ApplicationPresentationOptions)
	CurrentEvent() Event
	MainWindow() Window
	EffectiveAppearance() Appearance
	CurrentSystemPresentationOptions() ApplicationPresentationOptions
	ApplicationIconImage() Image
	SetApplicationIconImage(value IImage)
	IsActive() bool
	IsFullKeyboardAccessEnabled() bool
	Delegate() ApplicationDelegateWrapper
	SetDelegate(value PApplicationDelegate)
	SetDelegateObject(valueObject objc.IObject)
	ServicesProvider() objc.Object
	SetServicesProvider(value objc.IObject)
	UserInterfaceLayoutDirection() UserInterfaceLayoutDirection
	OrderedDocuments() []Document
	OcclusionState() ApplicationOcclusionState
	Windows() []Window
	HelpMenu() Menu
	SetHelpMenu(value IMenu)
	OrderedWindows() []Window
	WindowsMenu() Menu
	SetWindowsMenu(value IMenu)
	DockTile() DockTile
	IsAutomaticCustomizeTouchBarMenuItemEnabled() bool
	SetAutomaticCustomizeTouchBarMenuItemEnabled(value bool)
	MainMenu() Menu
	SetMainMenu(value IMenu)
	ModalWindow() Window
	EnabledRemoteNotificationTypes() RemoteNotificationType
	Appearance() Appearance
	SetAppearance(value IAppearance)
	IsProtectedDataAvailable() bool
}

// An object that manages an app’s main event loop and resources used by all of that app’s objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication?language=objc
type Application struct {
	Responder
}

func ApplicationFrom(ptr unsafe.Pointer) Application {
	return Application{
		Responder: ResponderFrom(ptr),
	}
}

func (ac _ApplicationClass) Alloc() Application {
	rv := objc.Call[Application](ac, objc.Sel("alloc"))
	return rv
}

func Application_Alloc() Application {
	return ApplicationClass.Alloc()
}

func (ac _ApplicationClass) New() Application {
	rv := objc.Call[Application](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewApplication() Application {
	return ApplicationClass.New()
}

func (a_ Application) Init() Application {
	rv := objc.Call[Application](a_, objc.Sel("init"))
	return rv
}

// Sends the given action message to the given target. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428509-sendaction?language=objc
func (a_ Application) SendActionToFrom(action objc.Selector, target objc.IObject, sender objc.IObject) bool {
	rv := objc.Call[bool](a_, objc.Sel("sendAction:to:from:"), action, target, sender)
	return rv
}

// Finishes a modal session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428438-endmodalsession?language=objc
func (a_ Application) EndModalSession(session ModalSession) {
	objc.Call[objc.Void](a_, objc.Sel("endModalSession:"), session)
}

// Starts a user attention request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428358-requestuserattention?language=objc
func (a_ Application) RequestUserAttention(requestType RequestUserAttentionType) int {
	rv := objc.Call[int](a_, objc.Sel("requestUserAttention:"), requestType)
	return rv
}

// Cancels a previous user attention request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428683-canceluserattentionrequest?language=objc
func (a_ Application) CancelUserAttentionRequest(request int) {
	objc.Call[objc.Void](a_, objc.Sel("cancelUserAttentionRequest:"), request)
}

// Dispatches an event to other objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428359-sendevent?language=objc
func (a_ Application) SendEvent(event IEvent) {
	objc.Call[objc.Void](a_, objc.Sel("sendEvent:"), objc.Ptr(event))
}

// Updates the Window menu item for a given window to reflect the edited status of that window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428422-updatewindowsitem?language=objc
func (a_ Application) UpdateWindowsItem(win IWindow) {
	objc.Call[objc.Void](a_, objc.Sel("updateWindowsItem:"), objc.Ptr(win))
}

// Arranges windows listed in the Window menu in front of all other windows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428739-arrangeinfront?language=objc
func (a_ Application) ArrangeInFront(sender objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("arrangeInFront:"), sender)
}

// Removes the Window menu item for a given window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428625-removewindowsitem?language=objc
func (a_ Application) RemoveWindowsItem(win IWindow) {
	objc.Call[objc.Void](a_, objc.Sel("removeWindowsItem:"), objc.Ptr(win))
}

// Hides all the receiver’s windows, and the next app in line is activated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428733-hide?language=objc
func (a_ Application) Hide(sender objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("hide:"), sender)
}

// Registers the pasteboard types the receiver can send and receive in response to service requests. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428751-registerservicesmenusendtypes?language=objc
func (a_ Application) RegisterServicesMenuSendTypesReturnTypes(sendTypes []PasteboardType, returnTypes []PasteboardType) {
	objc.Call[objc.Void](a_, objc.Sel("registerServicesMenuSendTypes:returnTypes:"), sendTypes, returnTypes)
}

// Searches for an object that can receive the message specified by the given selector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428658-targetforaction?language=objc
func (a_ Application) TargetForActionToFrom(action objc.Selector, target objc.IObject, sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("targetForAction:to:from:"), action, target, sender)
	return rv
}

// Searches for the string in the user interface. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1420808-searchstring?language=objc
func (a_ Application) SearchStringInUserInterfaceItemStringSearchRangeFoundRange(searchString string, stringToSearch string, searchRange foundation.Range, foundRange *foundation.Range) bool {
	rv := objc.Call[bool](a_, objc.Sel("searchString:inUserInterfaceItemString:searchRange:foundRange:"), searchString, stringToSearch, searchRange, foundRange)
	return rv
}

// Stops a modal event loop, allowing you to return a custom result code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428748-stopmodalwithcode?language=objc
func (a_ Application) StopModalWithCode(returnCode ModalResponse) {
	objc.Call[objc.Void](a_, objc.Sel("stopModalWithCode:"), returnCode)
}

// Unregister for notifications received from Apple Push Notification service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428747-unregisterforremotenotifications?language=objc
func (a_ Application) UnregisterForRemoteNotifications() {
	objc.Call[objc.Void](a_, objc.Sel("unregisterForRemoteNotifications"))
}

// Returns the window corresponding to the specified window number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428464-windowwithwindownumber?language=objc
func (a_ Application) WindowWithWindowNumber(windowNum int) Window {
	rv := objc.Call[Window](a_, objc.Sel("windowWithWindowNumber:"), windowNum)
	return rv
}

// Runs a given modal session, as defined in a previous invocation of beginModalSessionForWindow:relativeToWindow:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428590-runmodalsession?language=objc
func (a_ Application) RunModalSession(session ModalSession) ModalResponse {
	rv := objc.Call[ModalResponse](a_, objc.Sel("runModalSession:"), session)
	return rv
}

// Stops a modal event loop. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428489-stopmodal?language=objc
func (a_ Application) StopModal() {
	objc.Call[objc.Void](a_, objc.Sel("stopModal"))
}

// Aborts the event loop started by runModalForWindow:relativeToWindow: or runModalSession:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428753-abortmodal?language=objc
func (a_ Application) AbortModal() {
	objc.Call[objc.Void](a_, objc.Sel("abortModal"))
}

// Allows an app to extend its state restoration period. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1526248-extendstaterestoration?language=objc
func (a_ Application) ExtendStateRestoration() {
	objc.Call[objc.Void](a_, objc.Sel("extendStateRestoration"))
}

// Disables relaunching the app on login. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428376-disablerelaunchonlogin?language=objc
func (a_ Application) DisableRelaunchOnLogin() {
	objc.Call[objc.Void](a_, objc.Sel("disableRelaunchOnLogin"))
}

// Executes a block for each of the app's windows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1644472-enumeratewindowswithoptions?language=objc
func (a_ Application) EnumerateWindowsWithOptionsUsingBlock(options WindowListOptions, block func(window Window, stop *bool)) {
	objc.Call[objc.Void](a_, objc.Sel("enumerateWindowsWithOptions:usingBlock:"), options, block)
}

// Displays a standard About window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428724-orderfrontstandardaboutpanel?language=objc
func (a_ Application) OrderFrontStandardAboutPanel(sender objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("orderFrontStandardAboutPanel:"), sender)
}

// Completes the extended state restoration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1526245-completestaterestoration?language=objc
func (a_ Application) CompleteStateRestoration() {
	objc.Call[objc.Void](a_, objc.Sel("completeStateRestoration"))
}

// Opens the character palette. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428455-orderfrontcharacterpalette?language=objc
func (a_ Application) OrderFrontCharacterPalette(sender objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("orderFrontCharacterPalette:"), sender)
}

// Removes all events matching the given mask and generated before the specified event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428652-discardeventsmatchingmask?language=objc
func (a_ Application) DiscardEventsMatchingMaskBeforeEvent(mask EventMask, lastEvent IEvent) {
	objc.Call[objc.Void](a_, objc.Sel("discardEventsMatchingMask:beforeEvent:"), mask, objc.Ptr(lastEvent))
}

// Sets whether the receiver’s windows need updating when the receiver has finished processing the current event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428451-setwindowsneedupdate?language=objc
func (a_ Application) SetWindowsNeedUpdate(needUpdate bool) {
	objc.Call[objc.Void](a_, objc.Sel("setWindowsNeedUpdate:"), needUpdate)
}

// Returns the app’s activation policy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428703-activationpolicy?language=objc
func (a_ Application) ActivationPolicy() ApplicationActivationPolicy {
	rv := objc.Call[ApplicationActivationPolicy](a_, objc.Sel("activationPolicy"))
	return rv
}

// Show or hides the interface for customizing the Touch Bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/2646920-toggletouchbarcustomizationpalet?language=objc
func (a_ Application) ToggleTouchBarCustomizationPalette(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("toggleTouchBarCustomizationPalette:"), sender)
	return rv
}

// Activates the app, opens any files specified by the NSOpen user default, and unhighlights the app’s icon. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428771-finishlaunching?language=objc
func (a_ Application) FinishLaunching() {
	objc.Call[objc.Void](a_, objc.Sel("finishLaunching"))
}

// Suppresses the usual window ordering in handling the most recent mouse-down event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428640-preventwindowordering?language=objc
func (a_ Application) PreventWindowOrdering() {
	objc.Call[objc.Void](a_, objc.Sel("preventWindowOrdering"))
}

// Enables relaunching the app on login. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428453-enablerelaunchonlogin?language=objc
func (a_ Application) EnableRelaunchOnLogin() {
	objc.Call[objc.Void](a_, objc.Sel("enableRelaunchOnLogin"))
}

// Creates and executes a new thread based on the specified target and selector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428374-detachdrawingthread?language=objc
func (ac _ApplicationClass) DetachDrawingThreadToTargetWithObject(selector objc.Selector, target objc.IObject, argument objc.IObject) {
	objc.Call[objc.Void](ac, objc.Sel("detachDrawingThread:toTarget:withObject:"), selector, target, argument)
}

// Creates and executes a new thread based on the specified target and selector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428374-detachdrawingthread?language=objc
func Application_DetachDrawingThreadToTargetWithObject(selector objc.Selector, target objc.IObject, argument objc.IObject) {
	ApplicationClass.DetachDrawingThreadToTargetWithObject(selector, target, argument)
}

// Restores hidden windows without activating their owner (the receiver). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428566-unhidewithoutactivation?language=objc
func (a_ Application) UnhideWithoutActivation() {
	objc.Call[objc.Void](a_, objc.Sel("unhideWithoutActivation"))
}

// Unregister an object that provides help data to your app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1420820-unregisteruserinterfaceitemsearc?language=objc
func (a_ Application) UnregisterUserInterfaceItemSearchHandler(handler PUserInterfaceItemSearching) {
	po0 := objc.WrapAsProtocol("NSUserInterfaceItemSearching", handler)
	objc.Call[objc.Void](a_, objc.Sel("unregisterUserInterfaceItemSearchHandler:"), po0)
}

// Unregister an object that provides help data to your app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1420820-unregisteruserinterfaceitemsearc?language=objc
func (a_ Application) UnregisterUserInterfaceItemSearchHandlerObject(handlerObject objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("unregisterUserInterfaceItemSearchHandler:"), objc.Ptr(handlerObject))
}

// Displays the receiver’s page layout panel, an instance of NSPageLayout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1397808-runpagelayout?language=objc
func (a_ Application) RunPageLayout(sender objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("runPageLayout:"), sender)
}

// Logs a given exception by calling NSLog(). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428396-reportexception?language=objc
func (a_ Application) ReportException(exception foundation.IException) {
	objc.Call[objc.Void](a_, objc.Sel("reportException:"), objc.Ptr(exception))
}

// Brings up the color panel, an instance of NSColorPanel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1534097-orderfrontcolorpanel?language=objc
func (a_ Application) OrderFrontColorPanel(sender objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("orderFrontColorPanel:"), sender)
}

// Deactivates the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428428-deactivate?language=objc
func (a_ Application) Deactivate() {
	objc.Call[objc.Void](a_, objc.Sel("deactivate"))
}

// Adds an item to the Window menu for a given window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428660-addwindowsitem?language=objc
func (a_ Application) AddWindowsItemTitleFilename(win IWindow, string_ string, isFilename bool) {
	objc.Call[objc.Void](a_, objc.Sel("addWindowsItem:title:filename:"), objc.Ptr(win), string_, isFilename)
}

// Sends an update message to each onscreen window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428675-updatewindows?language=objc
func (a_ Application) UpdateWindows() {
	objc.Call[objc.Void](a_, objc.Sel("updateWindows"))
}

// Register for notifications sent by Apple Push Notification service (APNs). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/2967172-registerforremotenotifications?language=objc
func (a_ Application) RegisterForRemoteNotifications() {
	objc.Call[objc.Void](a_, objc.Sel("registerForRemoteNotifications"))
}

// Adds a given event to the receiver’s event queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428710-postevent?language=objc
func (a_ Application) PostEventAtStart(event IEvent, flag bool) {
	objc.Call[objc.Void](a_, objc.Sel("postEvent:atStart:"), objc.Ptr(event), flag)
}

// Unhides all apps, including the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428737-unhideallapplications?language=objc
func (a_ Application) UnhideAllApplications(sender objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("unhideAllApplications:"), sender)
}

// Places the receiver in context-sensitive help mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1500925-activatecontexthelpmode?language=objc
func (a_ Application) ActivateContextHelpMode(sender objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("activateContextHelpMode:"), sender)
}

// Attempts to modify the app’s activation policy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428621-setactivationpolicy?language=objc
func (a_ Application) SetActivationPolicy(activationPolicy ApplicationActivationPolicy) bool {
	rv := objc.Call[bool](a_, objc.Sel("setActivationPolicy:"), activationPolicy)
	return rv
}

// Miniaturizes all the receiver’s windows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428393-miniaturizeall?language=objc
func (a_ Application) MiniaturizeAll(sender objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("miniaturizeAll:"), sender)
}

// Returns the next event matching a given mask, or nil if no such event is found before a specified expiration date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428485-nexteventmatchingmask?language=objc
func (a_ Application) NextEventMatchingMaskUntilDateInModeDequeue(mask EventMask, expiration foundation.IDate, mode foundation.RunLoopMode, deqFlag bool) Event {
	rv := objc.Call[Event](a_, objc.Sel("nextEventMatchingMask:untilDate:inMode:dequeue:"), mask, objc.Ptr(expiration), mode, deqFlag)
	return rv
}

// Hides all apps, except the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428746-hideotherapplications?language=objc
func (a_ Application) HideOtherApplications(sender objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("hideOtherApplications:"), sender)
}

// Handles errors that might occur when the user attempts to open or print files. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428592-replytoopenorprint?language=objc
func (a_ Application) ReplyToOpenOrPrint(reply ApplicationDelegateReply) {
	objc.Call[objc.Void](a_, objc.Sel("replyToOpenOrPrint:"), reply)
}

// If your project is properly registered, and the necessary keys have been set in the property list, this method launches Help Viewer and displays the first page of your app’s help book. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1500910-showhelp?language=objc
func (a_ Application) ShowHelp(sender objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("showHelp:"), sender)
}

// Responds to NSTerminateLater once the app knows whether it can terminate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428594-replytoapplicationshouldterminat?language=objc
func (a_ Application) ReplyToApplicationShouldTerminate(shouldTerminate bool) {
	objc.Call[objc.Void](a_, objc.Sel("replyToApplicationShouldTerminate:"), shouldTerminate)
}

// Starts the main event loop. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428631-run?language=objc
func (a_ Application) Run() {
	objc.Call[objc.Void](a_, objc.Sel("run"))
}

// Stops the main event loop. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428473-stop?language=objc
func (a_ Application) Stop(sender objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("stop:"), sender)
}

// Restores hidden windows to the screen and makes the receiver active. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428761-unhide?language=objc
func (a_ Application) Unhide(sender objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("unhide:"), sender)
}

// Register an object that provides help data to your app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1420818-registeruserinterfaceitemsearchh?language=objc
func (a_ Application) RegisterUserInterfaceItemSearchHandler(handler PUserInterfaceItemSearching) {
	po0 := objc.WrapAsProtocol("NSUserInterfaceItemSearching", handler)
	objc.Call[objc.Void](a_, objc.Sel("registerUserInterfaceItemSearchHandler:"), po0)
}

// Register an object that provides help data to your app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1420818-registeruserinterfaceitemsearchh?language=objc
func (a_ Application) RegisterUserInterfaceItemSearchHandlerObject(handlerObject objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("registerUserInterfaceItemSearchHandler:"), objc.Ptr(handlerObject))
}

// Changes the item for a given window in the Window menu to a given string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428689-changewindowsitem?language=objc
func (a_ Application) ChangeWindowsItemTitleFilename(win IWindow, string_ string, isFilename bool) {
	objc.Call[objc.Void](a_, objc.Sel("changeWindowsItem:title:filename:"), objc.Ptr(win), string_, isFilename)
}

// Terminates the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428417-terminate?language=objc
func (a_ Application) Terminate(sender objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("terminate:"), sender)
}

// Invoked to request that a window be restored. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1526233-restorewindowwithidentifier?language=objc
func (a_ Application) RestoreWindowWithIdentifierStateCompletionHandler(identifier UserInterfaceItemIdentifier, state foundation.ICoder, completionHandler func(arg0 Window, arg1 foundation.Error)) bool {
	rv := objc.Call[bool](a_, objc.Sel("restoreWindowWithIdentifier:state:completionHandler:"), identifier, objc.Ptr(state), completionHandler)
	return rv
}

// Displays a standard About window with information from a given options dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428479-orderfrontstandardaboutpanelwith?language=objc
func (a_ Application) OrderFrontStandardAboutPanelWithOptions(optionsDictionary map[AboutPanelOptionKey]objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("orderFrontStandardAboutPanelWithOptions:"), optionsDictionary)
}

// A Boolean value indicating whether the app is registered with Apple Push Notification service (APNs). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/2967173-registeredforremotenotifications?language=objc
func (a_ Application) IsRegisteredForRemoteNotifications() bool {
	rv := objc.Call[bool](a_, objc.Sel("isRegisteredForRemoteNotifications"))
	return rv
}

// The app’s Services menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428608-servicesmenu?language=objc
func (a_ Application) ServicesMenu() Menu {
	rv := objc.Call[Menu](a_, objc.Sel("servicesMenu"))
	return rv
}

// The app’s Services menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428608-servicesmenu?language=objc
func (a_ Application) SetServicesMenu(value IMenu) {
	objc.Call[objc.Void](a_, objc.Sel("setServicesMenu:"), objc.Ptr(value))
}

// A Boolean value indicating whether the main event loop is running. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428759-running?language=objc
func (a_ Application) IsRunning() bool {
	rv := objc.Call[bool](a_, objc.Sel("isRunning"))
	return rv
}

// A Boolean value indicating whether the app is hidden. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428416-hidden?language=objc
func (a_ Application) IsHidden() bool {
	rv := objc.Call[bool](a_, objc.Sel("isHidden"))
	return rv
}

// The window that currently receives keyboard events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428406-keywindow?language=objc
func (a_ Application) KeyWindow() Window {
	rv := objc.Call[Window](a_, objc.Sel("keyWindow"))
	return rv
}

// The presentation options that should be in effect for the system when this app is active. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428664-presentationoptions?language=objc
func (a_ Application) PresentationOptions() ApplicationPresentationOptions {
	rv := objc.Call[ApplicationPresentationOptions](a_, objc.Sel("presentationOptions"))
	return rv
}

// The presentation options that should be in effect for the system when this app is active. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428664-presentationoptions?language=objc
func (a_ Application) SetPresentationOptions(value ApplicationPresentationOptions) {
	objc.Call[objc.Void](a_, objc.Sel("setPresentationOptions:"), value)
}

// The last event object that the app retrieved from the event queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428668-currentevent?language=objc
func (a_ Application) CurrentEvent() Event {
	rv := objc.Call[Event](a_, objc.Sel("currentEvent"))
	return rv
}

// The app’s main window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428723-mainwindow?language=objc
func (a_ Application) MainWindow() Window {
	rv := objc.Call[Window](a_, objc.Sel("mainWindow"))
	return rv
}

// The appearance that AppKit uses to draw the app’s interface. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/2967171-effectiveappearance?language=objc
func (a_ Application) EffectiveAppearance() Appearance {
	rv := objc.Call[Appearance](a_, objc.Sel("effectiveAppearance"))
	return rv
}

// The set of app presentation options that are currently in effect for the system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428717-currentsystempresentationoptions?language=objc
func (a_ Application) CurrentSystemPresentationOptions() ApplicationPresentationOptions {
	rv := objc.Call[ApplicationPresentationOptions](a_, objc.Sel("currentSystemPresentationOptions"))
	return rv
}

// The image used for the app’s icon. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428744-applicationiconimage?language=objc
func (a_ Application) ApplicationIconImage() Image {
	rv := objc.Call[Image](a_, objc.Sel("applicationIconImage"))
	return rv
}

// The image used for the app’s icon. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428744-applicationiconimage?language=objc
func (a_ Application) SetApplicationIconImage(value IImage) {
	objc.Call[objc.Void](a_, objc.Sel("setApplicationIconImage:"), objc.Ptr(value))
}

// A Boolean value indicating whether this is the active app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428493-active?language=objc
func (a_ Application) IsActive() bool {
	rv := objc.Call[bool](a_, objc.Sel("isActive"))
	return rv
}

// A Boolean value indicating whether Full Keyboard Access is enabled in the Keyboard preference pane. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428469-fullkeyboardaccessenabled?language=objc
func (a_ Application) IsFullKeyboardAccessEnabled() bool {
	rv := objc.Call[bool](a_, objc.Sel("isFullKeyboardAccessEnabled"))
	return rv
}

// The app delegate object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428705-delegate?language=objc
func (a_ Application) Delegate() ApplicationDelegateWrapper {
	rv := objc.Call[ApplicationDelegateWrapper](a_, objc.Sel("delegate"))
	return rv
}

// The app delegate object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428705-delegate?language=objc
func (a_ Application) SetDelegate(value PApplicationDelegate) {
	po0 := objc.WrapAsProtocol("NSApplicationDelegate", value)
	objc.SetAssociatedObject(a_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](a_, objc.Sel("setDelegate:"), po0)
}

// The app delegate object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428705-delegate?language=objc
func (a_ Application) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// The object that provides the services the current app advertises in the Services menu of other apps. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428467-servicesprovider?language=objc
func (a_ Application) ServicesProvider() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("servicesProvider"))
	return rv
}

// The object that provides the services the current app advertises in the Services menu of other apps. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428467-servicesprovider?language=objc
func (a_ Application) SetServicesProvider(value objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setServicesProvider:"), value)
}

// The layout direction of the user interface. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428556-userinterfacelayoutdirection?language=objc
func (a_ Application) UserInterfaceLayoutDirection() UserInterfaceLayoutDirection {
	rv := objc.Call[UserInterfaceLayoutDirection](a_, objc.Sel("userInterfaceLayoutDirection"))
	return rv
}

// An array of document objects arranged according to the front-to-back ordering of their associated windows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1494283-ordereddocuments?language=objc
func (a_ Application) OrderedDocuments() []Document {
	rv := objc.Call[[]Document](a_, objc.Sel("orderedDocuments"))
	return rv
}

// The occlusion state of the app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428656-occlusionstate?language=objc
func (a_ Application) OcclusionState() ApplicationOcclusionState {
	rv := objc.Call[ApplicationOcclusionState](a_, objc.Sel("occlusionState"))
	return rv
}

// An array of the app’s window objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428402-windows?language=objc
func (a_ Application) Windows() []Window {
	rv := objc.Call[[]Window](a_, objc.Sel("windows"))
	return rv
}

// The help menu used by the app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428644-helpmenu?language=objc
func (a_ Application) HelpMenu() Menu {
	rv := objc.Call[Menu](a_, objc.Sel("helpMenu"))
	return rv
}

// The help menu used by the app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428644-helpmenu?language=objc
func (a_ Application) SetHelpMenu(value IMenu) {
	objc.Call[objc.Void](a_, objc.Sel("setHelpMenu:"), objc.Ptr(value))
}

// Returns the application instance, creating it if it doesn’t exist yet. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428360-sharedapplication?language=objc
func (ac _ApplicationClass) SharedApplication() Application {
	rv := objc.Call[Application](ac, objc.Sel("sharedApplication"))
	return rv
}

// Returns the application instance, creating it if it doesn’t exist yet. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428360-sharedapplication?language=objc
func Application_SharedApplication() Application {
	return ApplicationClass.SharedApplication()
}

// An array of window objects arranged according to their front-to-back ordering on the screen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1494287-orderedwindows?language=objc
func (a_ Application) OrderedWindows() []Window {
	rv := objc.Call[[]Window](a_, objc.Sel("orderedWindows"))
	return rv
}

// The Window menu of the app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428547-windowsmenu?language=objc
func (a_ Application) WindowsMenu() Menu {
	rv := objc.Call[Menu](a_, objc.Sel("windowsMenu"))
	return rv
}

// The Window menu of the app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428547-windowsmenu?language=objc
func (a_ Application) SetWindowsMenu(value IMenu) {
	objc.Call[objc.Void](a_, objc.Sel("setWindowsMenu:"), objc.Ptr(value))
}

// The app’s Dock tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428671-docktile?language=objc
func (a_ Application) DockTile() DockTile {
	rv := objc.Call[DockTile](a_, objc.Sel("dockTile"))
	return rv
}

// A Boolean value indicating whether the main menu contains an item for customizing the contents of the Touch Bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/2646923-automaticcustomizetouchbarmenuit?language=objc
func (a_ Application) IsAutomaticCustomizeTouchBarMenuItemEnabled() bool {
	rv := objc.Call[bool](a_, objc.Sel("isAutomaticCustomizeTouchBarMenuItemEnabled"))
	return rv
}

// A Boolean value indicating whether the main menu contains an item for customizing the contents of the Touch Bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/2646923-automaticcustomizetouchbarmenuit?language=objc
func (a_ Application) SetAutomaticCustomizeTouchBarMenuItemEnabled(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setAutomaticCustomizeTouchBarMenuItemEnabled:"), value)
}

// The app’s main menu bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428634-mainmenu?language=objc
func (a_ Application) MainMenu() Menu {
	rv := objc.Call[Menu](a_, objc.Sel("mainMenu"))
	return rv
}

// The app’s main menu bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428634-mainmenu?language=objc
func (a_ Application) SetMainMenu(value IMenu) {
	objc.Call[objc.Void](a_, objc.Sel("setMainMenu:"), objc.Ptr(value))
}

// The modal window displayed by the app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428610-modalwindow?language=objc
func (a_ Application) ModalWindow() Window {
	rv := objc.Call[Window](a_, objc.Sel("modalWindow"))
	return rv
}

// The types of push notifications that the app accepts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/1428776-enabledremotenotificationtypes?language=objc
func (a_ Application) EnabledRemoteNotificationTypes() RemoteNotificationType {
	rv := objc.Call[RemoteNotificationType](a_, objc.Sel("enabledRemoteNotificationTypes"))
	return rv
}

// The appearance associated with the app’s windows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/2967170-appearance?language=objc
func (a_ Application) Appearance() Appearance {
	rv := objc.Call[Appearance](a_, objc.Sel("appearance"))
	return rv
}

// The appearance associated with the app’s windows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/2967170-appearance?language=objc
func (a_ Application) SetAppearance(value IAppearance) {
	objc.Call[objc.Void](a_, objc.Sel("setAppearance:"), objc.Ptr(value))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/3752992-protecteddataavailable?language=objc
func (a_ Application) IsProtectedDataAvailable() bool {
	rv := objc.Call[bool](a_, objc.Sel("isProtectedDataAvailable"))
	return rv
}