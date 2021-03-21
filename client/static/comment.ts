import { YozCommentComponent } from "./comment-component";
import { YozTimeLineComponent } from "./timeline-component";
import { deparam } from "./deparam";
import { startMeasuring } from "./measure";
class YozComment {
  private container!: HTMLElement;
  private commentComponent?: YozCommentComponent;
  private timelineComponent?: YozTimeLineComponent;

  constructor() {
    this.container = document.createElement("div");
    document.body.appendChild(this.container);
    let params = deparam(location.search.replace("?", ""));
    if (!this.container) {
      console.error("未设定渲染容器");
    }
    startMeasuring(params.origin);
    this.container.className = "yoz-comment-container";
    window.YozCommentConfig = {
      token: params.token,
      pageTitle: params.title,
      pageUrl: params.url,
    };
    this.loadTheme(params.theme || "light")
    this.commentComponent = new YozCommentComponent();
    this.container?.appendChild(this.commentComponent.element);
    this.timelineComponent = new YozTimeLineComponent();
    this.container?.appendChild(this.timelineComponent.element);
    this.timelineComponent.getList();
    let self = this;
    this.commentComponent.setEvent(function () {
      self.timelineComponent?.getList();
    });
  }

  private loadTheme(theme: string) {
    return new Promise((resolve) => {
      const link = document.createElement("link");
      link.rel = "stylesheet";
      link.setAttribute("crossorigin", "anonymous");
      link.onload = resolve;
      link.href = `./static/themes/${theme}.css`;
      document.head.appendChild(link);
    });
  }
}

new YozComment();
