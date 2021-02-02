import { KBCommentComponent } from "./comment-component";
import { KBTimeLineComponent } from "./timeline-component";
class KBComment {
  private container: HTMLElement | null;
  private commentComponent?: KBCommentComponent;
  private timelineComponent?: KBTimeLineComponent;
  private config: KBCommentConfig = {
    theme: "light",
    apiBase: "",
    token: "light",
  };

  constructor() {
    this.container = document.getElementById("kb-comment");
    if (this.container) {
      this.container.className = "kb-comment-container";
      this.config.theme = this.container.dataset.theme || "light";
      this.config.apiBase = this.container.dataset.api || "";
      this.config.token = this.container.dataset.token || location.pathname;
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
    } else {
      console.error("未设定渲染容器");
    }
  }

  private load(): Promise<any> {
    return Promise.all([this.loadTheme(), this.loadLibrary("http://pv.sohu.com/cityjson?ie=utf-8"), this.loadLibrary("https://cdn.jsdelivr.net/npm/axios@0.21.1/dist/axios.min.js"), this.loadLibrary("https://cdn.jsdelivr.net/npm/js-md5@0.7.3/build/md5.min.js")]);
  }

  private loadTheme() {
    return new Promise((resolve) => {
      const link = document.createElement("link");
      link.rel = "stylesheet";
      link.setAttribute("crossorigin", "anonymous");
      link.onload = resolve;
      link.href = `/themes/${this.config.theme}.css`;
      document.head.appendChild(link);
    });
  }

  private loadLibrary(url: string) {
    return new Promise((resolve) => {
      const script = document.createElement("script");
      script.src = url;
      script.onload = resolve;
      document.body.appendChild(script);
    });
  }
}

new KBComment();
