<script>
  import { page } from "$app/stores";

  const pages = [
    { title: "Home", route: "/" },
    { title: "Anime", route: "/anime" },
    { title: "Manga", route: "/manga" },
  ];

  let burgerActive = $state(false);
</script>

<div class="navbar is-spaced is-transparent">
  <div class="navbar-brand">
    <!-- website logo -->
    <div class="is-skeleton">OtakuDesk</div>
    <!-- hamburger menu icon -->
    <a
      href
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
      <div class="control has-icons-right">
        <input
          class="input is-rounded mx-3"
          type="text"
          placeholder="Search manga or anime..."
        />
        <!-- TODO: replace with search icon -->
        <div class="icon is-small is-right ml-3">&</div>
      </div>
    </div>

    <!-- account related menu and buttons -->
    <div class="navbar-end">
      <div class="buttons">
        <button class="button is-success">Log In</button>
        <button class="button is-light is-outlined">Sign Up</button>
      </div>
    </div>
  </div>
</div>
