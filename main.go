package main

/*
#cgo  CFLAGS: -x         objective-c
#cgo LDFLAGS: -framework Cocoa
#cgo LDFLAGS: -framework UserNotifications

#import <Cocoa/Cocoa.h>
#import <UserNotifications/UserNotifications.h>

NSMenuItem *getMenuItem(NSString *title, SEL action, NSString *keyEquivalent, NSEventModifierFlags keyEquivalentModifierMask) {
	id menuItem = [[[NSMenuItem alloc] initWithTitle:title action:action keyEquivalent:keyEquivalent] autorelease];
	[menuItem setKeyEquivalentModifierMask:keyEquivalentModifierMask];
	return menuItem;
}

NSMenuItem *getMenuItemOfSeparator(void) {
	return [NSMenuItem separatorItem];
}

void SetMenus(void) {
	id menus = @[
		@{
			@"title": [[NSProcessInfo processInfo] processName],
			@"items": @[
				getMenuItem([@"Quit " stringByAppendingString:[[NSProcessInfo processInfo] processName]], @selector(terminate:), @"q", NSEventModifierFlagCommand),
			],
		},
		@{
			@"title": @"Edit",
			@"items": @[
				getMenuItem(@"Undo",       @selector(undo:),      @"z", NSEventModifierFlagCommand),
				getMenuItem(@"Redo",       @selector(redo:),      @"Z", NSEventModifierFlagCommand),
				getMenuItemOfSeparator(),
				getMenuItem(@"Cut",        @selector(cut:),       @"x", NSEventModifierFlagCommand),
				getMenuItem(@"Copy",       @selector(copy:),      @"c", NSEventModifierFlagCommand),
				getMenuItem(@"Paste",      @selector(paste:),     @"v", NSEventModifierFlagCommand),
				getMenuItem(@"Select All", @selector(selectAll:), @"a", NSEventModifierFlagCommand),
			],
		},
		@{
			@"title": @"Window",
			@"items": @[
				getMenuItem(@"Enter Full Screen", @selector(toggleFullScreen:), @"f", NSEventModifierFlagCommand | NSEventModifierFlagControl),
			],
		},
	];

	id mainMenu = [[NSMenu new] autorelease];
	[NSApp setMainMenu:mainMenu];

	for (id menu in menus) {
		id mainMenuItem = [[NSMenuItem new] autorelease];
		[mainMenu addItem:mainMenuItem];

		id submenu = [[[NSMenu alloc] initWithTitle:menu[@"title"]] autorelease];
		[mainMenuItem setSubmenu:submenu];

		for (id item in menu[@"items"]) {
			[submenu addItem:item];
		}
	}
}

void RequestNotificationAuthorization(void) {
	UNAuthorizationOptions options = (
		UNAuthorizationOptionBadge |
		UNAuthorizationOptionSound |
		UNAuthorizationOptionAlert
	);

	[[UNUserNotificationCenter currentNotificationCenter] requestAuthorizationWithOptions:options completionHandler:^(BOOL granted, NSError *error) {
		//
	}];
}

void Notify(char *title, char *body) {
	UNMutableNotificationContent *content = [[UNMutableNotificationContent new] autorelease];
	content.title    = [NSString stringWithUTF8String:title];
	content.subtitle = @"app.metalife.co.jp";
	content.body     = [NSString stringWithUTF8String:body];

	id request = [UNNotificationRequest requestWithIdentifier:[[NSUUID UUID] UUIDString] content:content trigger:nil];
	[[UNUserNotificationCenter currentNotificationCenter] addNotificationRequest:request withCompletionHandler:nil];
}
*/
import "C"
import webview "github.com/webview/webview_go"

func main() {
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("MetaLife")
	w.SetSize(1280, 720, webview.HintNone)
	C.SetMenus()
	C.RequestNotificationAuthorization()

	w.Bind("setTitle", w.SetTitle)
	w.Bind("notify", func(title string, body string) {
		C.Notify(C.CString(title), C.CString(body))
	})
	w.Init(`
		const observer = new MutationObserver(() => {
			setTitle(document.title);
		});

		window.addEventListener('load', () => {
			setTitle(document.title);

			const title = document.querySelector('title');
			observer.observe(title, { childList: true });
		});

		window.Notification = class {
			static get permission() {
				return 'granted';
			}

			static async requestPermission() {
				return 'granted';
			}

			constructor(title, options = {}) {
				notify(title, options.body);
			}

			close() {
				//
			}
		};
	`)

	w.Navigate("https://app.metalife.co.jp/spaces")
	w.Run()
}
