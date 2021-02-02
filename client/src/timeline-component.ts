import { timeAgo } from "./time-ago";
export class KBTimeLineComponent {
  private element: HTMLDivElement;
  private loading: boolean = false;
  public constructor(private config: KBCommentConfig) {
    this.element = document.createElement("div");
    this.element.className = "kb-comment-list";
  }

  private getList(page: number = 1) {
    axios.get(this.config.apiBase + "/page?token=" + this.config.token).then((res: any) => {
      if (res.data.code == 200) {
        this.element.innerHTML = this.renderCommentItem(res.data.data.records);
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
