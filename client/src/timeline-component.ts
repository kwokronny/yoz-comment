import { KBCommentComponent } from "./comment-component";
import { timeAgo } from "./time-ago";
import { scheduleMeasure } from "./measure";
export class KBTimeLineComponent {
  public element: HTMLDivElement;
  private moreDOM: HTMLDivElement;
  private page: number = 1;
  private loading: boolean = false;
  private commentComp: KBCommentComponent;
  public constructor(private config: KBCommentConfig) {
    this.element = document.createElement("div");
    this.element.className = "kb-comment-list";
    this.commentComp = new KBCommentComponent(this.config);

    this.moreDOM = document.createElement("div");
    this.moreDOM.style.textAlign = "center";
    this.moreDOM.innerHTML = `<button class="more-btn">加载更多</button>`;
    this.moreDOM.querySelector(".more-btn")!.addEventListener("click", () => {
      this.page++;
      this.getList();
    });
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
        scheduleMeasure();
      }
    });
  }

  public getList(page?: number) {
    if (this.loading == true) return;
    this.loading = true;
    page = page || this.page || 1;
    // this.element.removeChild(this.moreDOM);
    if (page == 1) {
      this.element.innerHTML = "";
    }
    let loadingDOM = document.createElement("div");
    loadingDOM.innerHTML = `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" style="margin: auto; display: block; shape-rendering: auto;" width="80px" height="80px" viewBox="0 0 100 100" preserveAspectRatio="xMidYMid">
    <circle cx="30" cy="50" fill="#1d3f72" r="20">
      <animate attributeName="cx" repeatCount="indefinite" dur="1s" keyTimes="0;0.5;1" values="30;70;30" begin="-0.5s"></animate>
    </circle>
    <circle cx="70" cy="50" fill="#5699d2" r="20">
      <animate attributeName="cx" repeatCount="indefinite" dur="1s" keyTimes="0;0.5;1" values="30;70;30" begin="0s"></animate>
    </circle>
    <circle cx="30" cy="50" fill="#1d3f72" r="20">
      <animate attributeName="cx" repeatCount="indefinite" dur="1s" keyTimes="0;0.5;1" values="30;70;30" begin="-0.5s"></animate>
      <animate attributeName="fill-opacity" values="0;0;1;1" calcMode="discrete" keyTimes="0;0.499;0.5;1" dur="1s" repeatCount="indefinite"></animate>
    </circle>
    </svg>`;
    this.element.appendChild(loadingDOM);
    let params = {
      token: this.config.token,
      page,
    };
    axios
      .get(this.config.apiBase + "api/page", {
        params,
      })
      .then((res: any) => {
        this.loading = false;
        this.element.removeChild(loadingDOM);
        if (res.data.code == 200) {
          let data = res.data.data;
          let pageDOM = document.createElement("div");
          pageDOM.innerHTML = this.renderCommentItem(data.records);
          this.element.appendChild(pageDOM);
          if (data.page * data.pageSize < data.total) {
            this.element.appendChild(this.moreDOM);
          }
          scheduleMeasure();
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
