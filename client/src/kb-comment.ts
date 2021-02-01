import { timeAgo } from "./time-ago";
import { KBCommentComponent } from "./comment-component";
class KBComment {
  private container: HTMLElement | null;
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
        let form = new KBCommentComponent(this.config, this.getList);
        this.container?.appendChild(form.container);
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
      link.href = `/themes/${this.theme}.css`;
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

  private getList(page: number = 1) {
    var listContainer = document.createElement("div");
    listContainer.className = "kb-comment-list";
    axios.get(this.api + "/page?token=" + this.token).then((res: any) => {
      if (res.data.code == 200) {
        if (this.container) {
          listContainer.innerHTML = this.renderCommentItem(res.data.data.records);
          this.container.appendChild(listContainer);
        }
      }
    });
  }

  private renderCommentItem(list: CommentItem[]): string {
    return list.reduce((html: string, item: CommentItem) => {
      html += `
				<div class="comment-item">
					<div class="comment-avatar">
						<img src="https://s.gravatar.com/avatar/${md5(item.mail)}?s=50&d=retro&r=g" />
					</div>
					<div class="comment-message clear-right">
						<div>
							<div class="comment-time">${timeAgo(item.createdAt)}</div>
							<div class="comment-nickname"><a href="${item.site}">${item.nickName}</a></div>
						</div>
						<div class="comment-content">${item.content}</div>
					</div>
					<div class="comment-replys">
						${Array.isArray(item.replys) ? this.renderCommentItem(item.replys) : ""}
					</div>
				</div>`;
      return html;
    }, "");
  }
}

new KBComment();
