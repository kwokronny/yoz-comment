import { KBCommentComponent } from "./comment-component";
import { timeAgo } from "./time-ago";
export class KBTimeLineComponent {
  public element: HTMLDivElement;
  private loading: boolean = false;
  private commentComp: KBCommentComponent;
  public constructor(private config: KBCommentConfig) {
    this.element = document.createElement("div");
    this.element.className = "kb-comment-list";
    this.commentComp = new KBCommentComponent(this.config);
    this.element.addEventListener("click", (event: Event) => {
      let target: HTMLElement = event.target as HTMLElement;
      if (target.className == "reply-btn") {
        if (target.parentNode?.contains(this.commentComp.element)) {
          target.parentNode?.removeChild(this.commentComp.element);
        } else {
          target.parentNode?.appendChild(this.commentComp.element);
          this.commentComp.setReply(target.dataset.pid || "", target.dataset.rid || "");
          this.commentComp.setEvent(() => {
            target.parentNode?.removeChild(this.commentComp.element);
            this.getList();
          });
        }
      }
    });
  }

  public getList(page: number = 1) {
    if (this.loading == true) return;
    this.loading = true;
    this.element.innerHTML = `<div class="loading">loading</div>`;
    let params = {
      token: this.config.token,
      page,
    };
    axios
      .get(this.config.apiBase + "/page", {
        params,
      })
      .then((res: any) => {
        this.loading = false;
        if (res.data.code == 200) {
          this.element.innerHTML = this.renderCommentItem(res.data.data.records);
        }
      });
  }

  private renderCommentItem(list: KBCommentItem[], first: number = 0): string {
    return list.reduce((html: string, item: KBCommentItem) => {
      let pid = first == 0 ? item.id : first;
      html += `
				<div class="comment-item">
					<div class="comment-avatar">
						<img src="https://s.gravatar.com/avatar/${md5(item.mail)}?s=50&d=retro&r=g" />
					</div>
					<div class="comment-message clear-right">
						<div>
							<div class="comment-time">${timeAgo(item.createdAt)}</div>
							<div class="comment-nickname"><a target="_black" href="${item.site}">${item.nickName}</a></div>
						</div>
            <div class="comment-content">${item.content}</div>
            <div class="comment-option"><a class="reply-btn" data-rid="${item.id}" data-pid="${pid}" href="javascript:">回复</a></div>
					</div>
					<div class="comment-replys">
						${Array.isArray(item.replys) ? this.renderCommentItem(item.replys, pid) : ""}
					</div>
				</div>`;
      return html;
    }, "");
  }
}
