<script lang="ts">
  import { postApi } from "../service/api";
  import { auth } from "../stores/store.svelte";

  let email = $state("");
  let password = $state("");

  async function signup() {
    try {
      const response = await postApi({
        path: "/auth/signup",
        data: { email, password },
      });
      auth.token = response.token;
    } catch (error) {}
  }

  async function signin() {
    try {
      const response = await postApi({
        path: "/auth/signin",
        data: { email, password },
      });
      auth.token = response.token;
    } catch (error) {}
  }
</script>

<article>
  <input type="email" bind:value={email} />
  <input type="password" bind:value={password} />
  <div class="grid">
    <button class="outline" onclick={signup}>Sign Up</button>
    <button onclick={signin}>Sign In</button>
  </div>
</article>
