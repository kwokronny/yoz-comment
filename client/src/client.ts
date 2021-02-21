import { param } from "./deparam";
import { ResizeMessage } from './measure';
let script = document.currentScript as HTMLScriptElement;
if (script === undefined) {
  script = document.querySelector("script#KBComment") as HTMLScriptElement;
}

let attrs: Record<string, string> = {};
for (let i = 0; i < script.attributes.length; i++) {
  let attr = script.attributes.item(i)!;
  attrs[attr.name.replace(/^data-/, "")] = attr.value;
}
attrs.origin = location.origin;
document.head.insertAdjacentHTML(
  "afterbegin",
  `<style>
			.kbcomment {
				position: relative;
				box-sizing: border-box;
				width: 100%;
				max-width: 760px;
				margin-left: auto;
				margin-right: auto;
			}
			.kbcomment-frame {
				color-scheme: light;
				position: absolute;
				left: 0;
				right: 0;
				width: 1px;
				min-width: 100%;
				max-width: 100%;
				height: 100%;
				border: 0;
			}
		</style>`
);
let url = script.src.replace(/static\/client.js$/, "index.html");
script.insertAdjacentHTML(
  "afterend",
  `<div class="kbcomment">
    <iframe class="kbcomment-frame" title="Comments" scrolling="no" src="${url}?${param(attrs)}" loading="lazy"></iframe>
  </div>`
);
const container = script.nextElementSibling as HTMLDivElement;
script.parentElement!.removeChild(script);

// adjust the iframe's height when the height of it's content changes
addEventListener('message', event => {
	console.log(event)
  // if (event.origin !== ) {
  //   return;
  // }
  const data = event.data as ResizeMessage;
  if (data && data.type === 'resize' && data.height) {
    container.style.height = `${data.height}px`;
  }
});