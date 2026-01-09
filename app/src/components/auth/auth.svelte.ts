import { postApi } from "../../service/api";
import { auth } from "../../stores/store.svelte";

export async function signup(email: string, password: string) {
  try {
    const response = await postApi({
      path: "/auth/signup",
      data: { email, password },
    });
    auth.token = response.token;
  } catch (error) { }
}

export async function signin(email: string, password: string) {
  try {
    const response = await postApi({
      path: "/auth/signin",
      data: { email, password },
    });
    auth.token = response.token;
  } catch (error) { }
}