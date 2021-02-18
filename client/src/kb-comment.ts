import { KBCommentComponent } from "./comment-component";
import { KBTimeLineComponent } from "./timeline-component";
import { deparam } from "./deparam";
import { startMeasuring } from './measure';
class KBComment {
  private container!: HTMLElement;
  private commentComponent?: KBCommentComponent;
  private timelineComponent?: KBTimeLineComponent;
  private config: KBCommentConfig = {
    theme: "light",
    apiBase: "",
    token: "",
  };

  constructor() {
    this.container = document.createElement("div");
    document.body.appendChild(this.container);
    let params = deparam(location.search.replace("?", ""));
    if (!this.container) {
      console.error("未设定渲染容器");
    }
    startMeasuring(params.origin)
    this.container.className = "kb-comment-container";
    this.config.theme = params.theme || "light";
    this.config.apiBase = params.api || "";
    this.config.token = params.token || location.pathname;
    this.load().then(() => {
      this.commentComponent = new KBCommentComponent(this.config);
      this.container?.appendChild(this.commentComponent.element);
      this.timelineComponent = new KBTimeLineComponent(this.config);
      this.container?.appendChild(this.timelineComponent.element);
      this.timelineComponent.getList();
      let self = this;
      this.commentComponent.setEvent(function () {
        self.timelineComponent?.getList();
      });
    });
  }

  private load(): Promise<any> {
    return Promise.all([this.loadTheme()]);
  }

  private loadTheme() {
    return new Promise((resolve) => {
      const link = document.createElement("link");
      link.rel = "stylesheet";
      link.setAttribute("crossorigin", "anonymous");
      link.onload = resolve;
      link.href = `/web/themes/${this.config.theme}.css`;
      document.head.appendChild(link);
    });
  }
}

new KBComment();
