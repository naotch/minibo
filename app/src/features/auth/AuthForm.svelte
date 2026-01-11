<script lang="ts">
  import { signup, signin } from "./auth.svelte";

  let email = $state("");
  let password = $state("");
</script>

{#snippet button(title: string, type: string, onclick: () => void)}
  <button class={type} {onclick}>{title}</button>
{/snippet}

<article>
  <label for="email">メールアドレス</label>
  <input type="email" id="email" bind:value={email} />
  <label for="password">パスワード</label>
  <input type="password" id="password" bind:value={password} />
  <div class="grid">
    {@render button("新規作成", "outline", () => signup(email, password))}
    {@render button("サインイン", "primary", () => signin(email, password))}
  </div>
</article>
