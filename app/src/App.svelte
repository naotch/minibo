<script lang="ts">
  import LoginForm from "./components/UserForm.svelte";
  import { auth } from "./stores/store.svelte";
</script>

<main class="container">
  {#if !auth.token}
    <LoginForm />
  {:else}
    <p>Hello World!</p>
  {/if}
</main>
