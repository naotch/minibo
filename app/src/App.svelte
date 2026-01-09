<script lang="ts">
  import AuthForm from "./components/auth/SignupForm.svelte";
  import { auth } from "./stores/store.svelte";
</script>

<main class="container">
  {#if !auth.token}
    <AuthForm />
  {:else}
    <p>Hello World!</p>
  {/if}
</main>
