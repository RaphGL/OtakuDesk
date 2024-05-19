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
</script>

<svelte:window onkeyup={closeModalOnEscape} />

<div class="modal" class:is-active={isActive}>
  <div onclick={onclose} class="modal-background"></div>

  <div class="modal-content box p-6">
    <figure class="image is-3by1">
      <image src="#" alt="website logo"></image>
    </figure>
    <form>
      {#if mode === "register"}
        <div class="field">
          <div class="control has-icons-left">
            <input class="input" type="text" placeholder="username" />
            <span class="icon is-small is-left">
              <UserIcon />
            </span>
          </div>
        </div>
      {/if}

      <!-- email field -->
      <div class="field">
        <div class="control has-icons-left">
          <input class="input" type="email" placeholder="email" />
          <span class="icon is-small is-left">
            <EmailIcon />
          </span>
        </div>
      </div>
      <!-- password field -->
      <div class="field">
        <p class="control has-icons-left">
          <input class="input" type="password" placeholder="password" />
          <span class="icon is-small is-left">
            <PasswordIcon />
          </span>
        </p>
      </div>
      <!-- login and register buttons -->
      <div class="field is-grouped">
        {#if mode === "login"}
            <button class="button is-fullwidth is-success">Login</button>
            <a href="#" class="button is-text">Forgot your password?</a>
        {:else if mode === "register"}
            <button class="button is-fullwidth is-success">Register</button>
        {/if}
      </div>
    </form>
  </div>

  <div onclick={onclose} class="modal-close"></div>
</div>
