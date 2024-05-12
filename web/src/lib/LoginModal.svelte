<script>
  let { isActive, onclose } = $props();

  // close modal on escape key
  function closeModalOnEscape(e) {
    if (e.key === 'Escape') {
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
      <div class="field">
        <label for="username" class="label">Username</label>
        <div class="control">
          <input class="input" type="text" placeholder="username" />
        </div>
      </div>

      <!-- email field -->
      <div class="field">
        <label for="email" class="label">Email</label>
        <div class="control has-icons-left">
          <input class="input" type="email" placeholder="email" />
          <span class="icon is-small is-left">Email Icon</span>
        </div>
        <!-- password field -->
        <div class="field">
        <label for="password" class="label">Password</label>
          <p class="control has-icons-left">
            <input class="input" type="password" placeholder="password" />
            <span class="icon is-small is-left"> Password Icon </span>
          </p>
        </div>
      </div>
      <!-- login and register buttons -->
      <div class="field is-grouped">
        <div class="control">
          <button class="button is-success">Login</button>
        </div>
        <div class="control">
          <button class="button">Register</button>
        </div>
        <div class="control">
          <a href='#' class="button is-text">Forgot your password?</a>
        </div>
      </div>
    </form>
  </div>

  <div onclick={onclose} class="modal-close"></div>
</div>
