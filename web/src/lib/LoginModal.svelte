<script context="module" lang="ts">
  export type ModalMode = "register" | "login";
</script>

<script lang="ts">
  import EmailIcon from "$lib/icons/EmailIcon.svelte";
  import PasswordIcon from "$lib/icons/PasswordIcon.svelte";
  import UserIcon from "$lib/icons/UserIcon.svelte";

  type Props = {
    isActive: boolean;
    onclose: () => void;
    mode: ModalMode;
  };

  let { isActive = false, onclose, mode = "login" }: Props = $props();

  // close modal on escape key
  function closeModalOnEscape(e: KeyboardEvent) {
    if (e.key === "Escape") {
      onclose();
    }
  }

  let emailInput = $state(""),
    usernameInput = $state(""),
    passwordInput = $state(""),
    responseNotification = $state("");

  // todo: store jwt when register is successful
  async function register() {
    const payload = await fetch("http://localhost:8080/register", {
      mode: "cors",
      method: "POST",
      body: JSON.stringify({
        email: emailInput,
        username: usernameInput,
        password: passwordInput,
      }),
    });

    responseNotification = await payload.text();
  }

  // todo: add authorization jwt token from local storage to request header
  async function login() {
    const payload = await fetch("http://localhost:8080/login", {
      mode: "cors",
      method: "POST",
      body: JSON.stringify({
        username: usernameInput,
        password: passwordInput,
      }),
    });

    responseNotification = await payload.text();
  }
</script>

<svelte:window onkeyup={closeModalOnEscape} />

<div class="modal" class:is-active={isActive}>
  <div onclick={onclose} class="modal-background"></div>

  <div class="modal-content box p-6">
    <figure class="image is-3by1">
      <image src="#" alt="website logo"></image>
    </figure>
    <form>
      <div class="field">
        <div class="control has-icons-left">
          <input
            bind:value={usernameInput}
            class="input"
            type="text"
            placeholder="username"
          />
          <span class="icon is-small is-left">
            <UserIcon />
          </span>
        </div>
      </div>

      <!-- email field -->
      {#if mode === "register"}
        <div class="field">
          <div class="control has-icons-left">
            <input
              bind:value={emailInput}
              class="input"
              type="email"
              placeholder="email"
            />
            <span class="icon is-small is-left">
              <EmailIcon />
            </span>
          </div>
        </div>
      {/if}
      <!-- password field -->
      <div class="field">
        <p class="control has-icons-left">
          <input
            bind:value={passwordInput}
            class="input"
            type="password"
            placeholder="password"
          />
          <span class="icon is-small is-left">
            <PasswordIcon />
          </span>
        </p>
      </div>
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

      <!-- confirm successful login/registering -->
      {#if responseNotification != ""}
        <div class="notification has-text-centered has-text-weight-bold">
          {responseNotification}
        </div>
      {/if}
    </form>
  </div>

  <div onclick={onclose} class="modal-close"></div>
</div>
