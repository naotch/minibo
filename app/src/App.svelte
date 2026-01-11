<script lang="ts">
  import Header from "./features/shared/Header.svelte";
  import AuthForm from "./features/auth/AuthForm.svelte";
  import RecordForm from "./features/transaction/RecordForm.svelte";
  import Modal from "./features/shared/Modal.svelte";
  import { auth, modal } from "./stores/store.svelte";
</script>

<header class="container">
  <Header />
</header>
<main class="container">
  {#if modal.message}
    <Modal onclick={() => (modal.message = "")} />
  {/if}
  {#if !auth.token}
    <AuthForm />
  {:else}
    <RecordForm />
  {/if}
</main>
