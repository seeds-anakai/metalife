package main

/*
#cgo  CFLAGS: -x         objective-c
#cgo LDFLAGS: -framework Cocoa

#import <Cocoa/Cocoa.h>

NSMenuItem *getMenuItem(NSString *title, SEL action, NSString *keyEquivalent, NSEventModifierFlags keyEquivalentModifierMask) {
	NSMenuItem *menuItem = [[[NSMenuItem alloc] initWithTitle:title action:action keyEquivalent:keyEquivalent] autorelease];
	[menuItem setKeyEquivalentModifierMask:keyEquivalentModifierMask];
	return menuItem;
}

NSMenuItem *getMenuItemOfSeparator(void) {
	return [NSMenuItem separatorItem];
}

void SetMenu(void) {
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
*/
import "C"
import (
	webview "github.com/webview/webview_go"
)

func main() {
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("MetaLife")
	w.SetSize(1280, 720, webview.HintNone)
	C.SetMenu()

	w.Bind("setTitle", w.SetTitle)
	w.Init(`
		const observer = new MutationObserver(() => {
			setTitle(document.title);
		});

		window.addEventListener('load', () => {
			setTitle(document.title);

			const title = document.querySelector('title');
			observer.observe(title, { childList: true });
		});
	`)

	w.Navigate("https://app.metalife.co.jp/spaces")
	w.Run()
}
