import { postApi } from "../../service/api";
import { auth, modal } from "../../stores/store.svelte";

export async function signup(email: string, password: string) {
  try {
    const response = await postApi({
      path: "/auth/signup",
      data: { email, password },
    });
    auth.token = response.token;

  } catch (error: any) {
    const status = error.response?.status;
    modal.title = "エラー"
    switch (status) {
      case 400:
        modal.message = "メールアドレスまたはパスワードの入力形式が正しくありません"
        break;
      default:
        modal.message = "システムエラーが発生しました";
    }
  }
}

export async function signin(email: string, password: string) {
  try {
    const response = await postApi({
      path: "/auth/signin",
      data: { email, password },
    });
    auth.token = response.token;

  } catch (error: any) {
    const status = error.response?.status;
    modal.title = "エラー"
    switch (status) {
      case 400:
        modal.message = "メールアドレスまたはパスワードの入力形式が正しくありません"
        break;
      case 401:
        modal.message = "メールアドレスまたはパスワードが正しくありません"
        break;
      default:
        modal.message = "システムエラーが発生しました"
    }
  }
}