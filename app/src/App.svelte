<script lang="ts">
  import AuthForm from "./features/auth/AuthForm.svelte";
  import Modal from "./features/common/Modal.svelte";
  import { auth, modal } from "./stores/store.svelte";
</script>

<main class="container">
  {#if modal.message}
    <Modal onclick={() => (modal.message = "")} />
  {/if}

  {#if !auth.token}
    <AuthForm />
  {:else}
    <p>Hello World!</p>
  {/if}
</main>
