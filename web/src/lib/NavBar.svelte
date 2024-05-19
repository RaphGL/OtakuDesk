<script lang="ts">
  import LoginModal from "$lib/LoginModal.svelte";
  import { page } from "$app/stores";
  import SearchIcon from "$lib/icons/SearchIcon.svelte";
  import type { ModalMode } from "$lib/LoginModal.svelte";

  const pages = [
    { title: "Home", route: "/" },
    { title: "Anime", route: "/anime" },
    { title: "Manga", route: "/manga" },
  ];

  let burgerActive = $state(false);
  let modalActive = $state(false);
  // changes depending on if the user clicks the register or the login button
  let modalMode: ModalMode = $state("login");

  function showModal(mode: ModalMode) {
    modalMode = mode;
    modalActive = true;
  }

  function hideModal() {      
    modalActive = false;
  }
</script>

<div class="navbar is-spaced is-transparent">
  <div class="navbar-brand">
    <!-- website logo -->
    <div class="is-skeleton">OtakuDesk</div>
    <!-- hamburger menu icon -->
    <a
      href="#"
      onclick={() => (burgerActive = !burgerActive)}
      class:is-active={burgerActive}
      role="button"
      class="navbar-burger"
      aria-label="menu"
      aria-expanded="false"
    >
      <span aria-hidden="true"></span>
      <span aria-hidden="true"></span>
      <span aria-hidden="true"></span>
      <span aria-hidden="true"></span>
    </a>
  </div>

  <div class="navbar-menu" class:is-active={burgerActive}>
    <div class="navbar-start">
      <!-- link to website pages -->
      {#each pages as pageInfo}
        {@const isActivePage = $page.url.pathname === pageInfo.route}
        <a
          href={pageInfo.route}
          class="navbar-item"
          class:has-text-success={isActivePage}
          class:has-text-weight-bold={isActivePage}
        >
          {pageInfo.title}
        </a>
      {/each}

      <!-- search bar -->
      <div class="control has-icons-left mx-1">
        <div class="icon is-small is-left">
          <SearchIcon />
        </div>
        <input
          class="input is-rounded"
          type="text"
          placeholder="Search manga or anime..."
        />
        <!-- TODO: replace with search icon -->
      </div>
    </div>

    <!-- account related menu and buttons -->
    <div class="navbar-end">
      <div class="buttons my-3">
        <button onclick={() => showModal("login")} class="button is-success"
          >Log In</button
        >
        <button
          onclick={() => showModal("register")}
          class="button is-light is-outlined">Register</button
        >
      </div>
    </div>
  </div>
</div>

<LoginModal mode={modalMode} isActive={modalActive} onclose={hideModal} />
