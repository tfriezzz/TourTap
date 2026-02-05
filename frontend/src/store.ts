import { reactive } from "vue"
import type { AuthState } from "./types/auth_state"
import { readonly } from "vue"
import type User from 'types/user.ts'

const state = reactive<AuthState>({
  user: null,
  accesToken: null,
  refreshToken: null,
})


function setUser(user: User, accesToken: string, refreshToken: string) {
  state.user = user
  state.accesToken = accesToken
  state.refreshToken = refreshToken
}


function clearUser() {
  state.user = null
  state.accesToken = null
  state.refreshToken = null
}


export default {
  state: readonly(state),
  setUser,
  clearUser
}
