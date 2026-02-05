import { reactive } from "vue"
import type { AuthState } from "./types/auth_state"
import { readonly } from "vue"
import type { User } from "./types/user"

const state = reactive<AuthState>({
  user: null,
  accessToken: null,
  refreshToken: null,
})

const saved = localStorage.getItem('auth')
if (saved) {
  try {
    const parsed = JSON.parse(saved)
    state.user = parsed.user
    state.accessToken = parsed.accessToken
    state.refreshToken = parsed.refreshToken
  } catch (e) {
  }
}

function setUser(user: User, accessToken: string, refreshToken: string) {
  state.user = user
  state.accessToken = accessToken
  state.refreshToken = refreshToken

  localStorage.setItem('auth', JSON.stringify({
    user,
    accessToken,
    refreshToken
  }))
}


function clearUser() {
  state.user = null
  state.accessToken = null
  state.refreshToken = null
  localStorage.removeItem('auth')
}


export default {
  state: readonly(state),
  setUser,
  clearUser
}
