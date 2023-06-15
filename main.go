package main

import (
	"github.com/webview/webview"
)

func main() {
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("MetaLife")
	w.SetSize(1280, 720, webview.HintNone)

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
