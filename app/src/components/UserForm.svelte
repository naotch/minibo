<script lang="ts">
  import { postApi } from "../service/api";
  import { auth } from "../stores/store.svelte";

  let username = $state("");
  let password = $state("");

  async function signup() {
    try {
      const response = await postApi({
        path: "/signup",
        data: { username, password },
      });
    } catch (error) {}
  }
  async function signin() {
    try {
      const response = await postApi({
        path: "/signin",
        data: { username, password },
      });
      auth.token = response.token;
    } catch (error) {}
  }
</script>

<article>
  <input type="text" bind:value={username} />
  <input type="password" bind:value={password} />
  <div class="grid">
    <button class="outline" onclick={signup}>Sign Up</button>
    <button onclick={signin}>Sign In</button>
  </div>
</article>
