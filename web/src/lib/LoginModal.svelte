<script context="module" lang="ts">
  export type ModalMode = "register" | "login";
</script>

<script lang="ts">
  import EmailIcon from "$lib/icons/EmailIcon.svelte";
  import PasswordIcon from "$lib/icons/PasswordIcon.svelte";
  import UserIcon from "$lib/icons/UserIcon.svelte";
  import Input from "./helpers/Input.svelte";

  type Props = {
    isActive: boolean;
    onclose: () => void;
    mode: ModalMode;
  };

  let { isActive = false, onclose, mode = "login" }: Props = $props();

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

  async function register() {
    const resp = await fetch("http://localhost:8080/register", {
      mode: "cors",
      method: "POST",
      body: JSON.stringify({
        email: userInput.email,
        username: userInput.username,
        password: userInput.password,
      }),
    });

    responseNotification = await resp.text();
    if (resp.ok) {
      setTimeout(() => {
        closeModal();
      }, 1500);
    }
  }

  // todo: add authorization jwt token from local storage to request header
  async function login() {
    const resp = await fetch("http://localhost:8080/login", {
      mode: "cors",
      method: "POST",
      credentials: 'include',
      body: JSON.stringify({
        username: userInput.username,
        password: userInput.password,
      }),
    });

    if (resp.ok) {
      closeModal();
    }
  }
</script>

<svelte:window onkeyup={closeModalOnEscape} />

<div class="modal" class:is-active={isActive}>
  <div onclick={closeModal} class="modal-background"></div>
  <div class="modal-content box p-6">
    <figure class="image is-3by1">
      <image src="#" alt="website logo"></image>
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
      <div class="field is-grouped">
        {#if mode === "login"}
          <button class="button is-fullwidth is-success" onclick={login}
            >Login</button
          >
          <a href="#" class="button is-text">Forgot your password?</a>
        {:else if mode === "register"}
          <button class="button is-fullwidth is-success" onclick={register}
            >Register</button
          >
        {/if}
      </div>
      <!-- confirm successful registering -->
      {#if responseNotification != ""}
        <div
          class="notification is-info has-text-centered has-text-weight-bold"
        >
          {responseNotification}
        </div>
      {/if}
    </form>
  </div>

  <div onclick={closeModal} class="modal-close"></div>
</div>
