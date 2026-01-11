import { getApi } from "../../service/api";
import { auth, modal } from "../../stores/store.svelte";

export async function getTotal() {
  try {
    const res = await getApi({
      path: "/reports/total",
      token: auth.token
    })
    return res.total

  } catch (error: any) {
    const status = error.response?.status
    modal.title = "エラー"
    switch (status) {
      case 401:
        modal.message = "ログインしてください"
        break
      default:
        modal.message = "システムエラーが発生しました"
    }
  }
}