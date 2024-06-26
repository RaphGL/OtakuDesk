<script context="module" lang="ts">
  export type ModalMode = "register" | "login";
</script>

<script lang="ts">
  import EmailIcon from "$lib/icons/EmailIcon.svelte";
  import PasswordIcon from "$lib/icons/PasswordIcon.svelte";
  import UserIcon from "$lib/icons/UserIcon.svelte";
  import Input from "$lib/helpers/Input.svelte";
  import { getAuthStore } from "$lib/stores/auth.svelte";

  type Props = {
    isActive: boolean;
    onclose: () => void;
    mode: ModalMode;
  };

  let { isActive = false, onclose, mode = "login" }: Props = $props();
  const authStore = getAuthStore();

  let userInput = $state({
    email: "",
    username: "",
    password: "",
  });

  let responseNotification = $state("");

  function closeModal() {
    userInput = {
      email: "",
      username: "",
      password: "",
    };
    responseNotification = "";
    onclose();
  }

  // close modal on escape key
  function closeModalOnEscape(e: KeyboardEvent) {
    if (e.key === "Escape") {
      closeModal();
    }
  }

  async function login() {
    const resp = await authStore.login(userInput);
    responseNotification = resp.response;
    if (resp.ok) closeModal();
  }

  async function register() {
    const resp = await authStore.register(userInput);
    responseNotification = resp.response;
    if (resp.ok) {
      setTimeout(closeModal, 1500);
    }
  }
</script>

<svelte:window onkeyup={closeModalOnEscape} />

<div class="modal" class:is-active={isActive}>
  <div
    onkeyup={closeModal}
    role="button"
    tabindex="0"
    onclick={closeModal}
    class="modal-background"
  ></div>
  <div class="modal-content box p-6">
    <figure class="field image is-3by1">
      <img
        src="https://bulma.io/assets/images/placeholders/128x128.png"
        alt="website logo"
      />
    </figure>

    <form>
      <Input
        type="text"
        placeholder="username"
        bind:value={userInput.username}
        icon={UserIcon}
      />
      {#if mode === "register"}
        <Input
          type="email"
          placeholder="email"
          bind:value={userInput.email}
          icon={EmailIcon}
        />
      {/if}
      <Input
        type="password"
        placeholder="password"
        bind:value={userInput.password}
        icon={PasswordIcon}
      />
      <!-- login and register buttons -->
      {#if responseNotification === ""}
        <div class="field is-grouped">
          {#if mode === "login"}
            <button class="button is-fullwidth is-success" onclick={login}
              >Login</button
            >
            <button class="button is-text">Forgot your password?</button>
          {:else if mode === "register"}
            <button class="button is-fullwidth is-success" onclick={register}
              >Register</button
            >
          {/if}
        </div>
        <!-- confirm successful registering -->
      {:else}
        <div
          class="notification is-info has-text-centered has-text-weight-bold"
        >
          {responseNotification}
        </div>
      {/if}
    </form>
  </div>

  <button onclick={closeModal} class="modal-close"></button>
</div>
