import { getContext, setContext } from "svelte";

const STORE_NAME = 'auth_store';

type AuthResponse = {
  response: string,
  ok: boolean,
};

type LoginPayload = {
  username: string,
  password: string,
};

type RegisterPayload = LoginPayload & {
  email: string,
};

class AuthStore {
  #isLoggedIn = $state(false);
  #username = $state("");

  isLoggedIn(): boolean {
    return this.#isLoggedIn;
  }

  getUsername(): string {
    return this.#username;
  }

  // check if the session cookie is valid 
  async checkSessionValidity(): Promise<boolean> {
    const resp = await fetch("http://localhost:8080/is-auth", {
      credentials: "include",
      mode: "cors",
    });

    if (resp.ok) {
      const payload = await resp.json();
      this.#isLoggedIn = true;
      this.#username = payload.username;
    }

    return resp.ok;
  }

  async logout() {
    let resp = await fetch("http://localhost:8080/logout", {
      credentials: "include",
      mode: "cors",
      method: "POST",
    });

    if (resp.ok) {
      this.#isLoggedIn = false;
    }
  }

  async register(payload: RegisterPayload): Promise<AuthResponse> {
    const resp = await fetch("http://localhost:8080/register", {
      mode: "cors",
      method: "POST",
      body: JSON.stringify({
        email: payload.email,
        username: payload.username,
        password: payload.password,
      }),
    });

    return {
      response: await resp.text(),
      ok: resp.ok,
    };
  }

  async login(payload: LoginPayload): Promise<AuthResponse> {
    const resp = await fetch("http://localhost:8080/login", {
      mode: "cors",
      method: "POST",
      credentials: 'include',
      body: JSON.stringify({
        username: payload.username,
        password: payload.password,
      }),
    });

    if (resp.ok) {
      this.#username = payload.username;
      this.#isLoggedIn = true;
    }

    return {
      response: await resp.text(),
      ok: resp.ok,
    };
  }
}

export function createAuthStore(): AuthStore {
  const store = new AuthStore();
  return setContext(STORE_NAME, store);
}

export function getAuthStore(): AuthStore {
  return getContext(STORE_NAME);
}
