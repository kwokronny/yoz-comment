export class KBCommentComponent {
  public element: HTMLFormElement;
  private parentIdField: HTMLInputElement;
  private rIdField: HTMLInputElement;
  private nickNameField: HTMLInputElement;
  private mailField: HTMLInputElement;
  private siteField: HTMLInputElement;
  private contentField: HTMLTextAreaElement;
  private submitBtn: HTMLButtonElement;
  private submitting: boolean = false;
  private success: Function = () => {};

  public constructor(private config: KBCommentConfig) {
    this.config = config;
    this.element = document.createElement("form");
    this.element.className = "kb-comment-form";
    this.element.action = "javascript:";
    this.element.innerHTML = `
		<div class="user-info clearfix">
			<div class="is-reply">
				回复 用户名 <a class="reset-reply" href="javascript:">取消</a>
				<input type="hidden" name="parentId" value="0" />
				<input type="hidden" name="rId" value="0" />
			</div>
			<div class="input-col">
				<input type="text" name="nickname" maxlength="40" placeholder="昵称(必填)" required/>
			</div>
			<div class="input-col">
				<input type="email" name="mail" placeholder="邮箱(必填)" required/>
			</div>
			<div class="input-col">
				<input type="url" name="site" maxlength="40" placeholder="网址" />
			</div>
		</div>
    <div class="message">
      <textarea row="6" name="content" placeholder="请输入你的留言" required></textarea>
		</div>
		<div class="btn-group">
			<button type="submit">评论</button>
		</div>`;
    this.parentIdField = this.element.querySelector("input[name=parentId]") as HTMLInputElement;
    this.rIdField = this.element.querySelector("input[name=rId]") as HTMLInputElement;
    this.nickNameField = this.element.querySelector("input[name=nickname]") as HTMLInputElement;
    this.mailField = this.element.querySelector("input[name=mail]") as HTMLInputElement;
    this.siteField = this.element.querySelector("input[name=site]") as HTMLInputElement;
    this.contentField = this.element.querySelector("textarea") as HTMLTextAreaElement;
    this.submitBtn = this.element.querySelector("button[type=submit]") as HTMLButtonElement;
    this.element.querySelector("a.reset-reply")?.addEventListener("click", this.resetReply.bind(this));
    this.element.addEventListener("submit", this.onSubmitComment.bind(this), false);
  }

  public setEvent(successFunc: Function) {
    this.success = successFunc;
  }

  private getModel() {
    return {
      parentId: this.parentIdField.value,
      rId: this.rIdField.value,
      nickName: this.nickNameField.value,
      mail: this.mailField.value,
      site: this.siteField.value,
      content: this.contentField.value,
      ip: returnCitySN.cip,
      token: this.config.token,
    };
  }

  private onSubmitComment(event: Event) {
    event.preventDefault();
    if (this.submitting) {
      return;
    }
    this.submitting = true;
    this.submitBtn.disabled = true;
    axios.post(this.config.apiBase + "/comment", this.getModel(), (res: any) => {
      if (res.data.code == 200) {
        this.success();
      }
    });
    return false;
  }

  private resetReply() {
    this.setReply(0, 0);
  }

  public setReply(parentId: number, rId: number) {
    this.rIdField.value = rId.toString();
    this.parentIdField.value = parentId.toString();
  }
}
