<template>
  <center class="signup">
    <card class="signup-card box-shadow p18">
      <h2>Sign Up</h2>
      <el-form
        :model="model"
        :rules="rules"
        ref="form"
        @submit.native.prevent="onSingup"
      >
        <el-form-item prop="email">
          <el-input
            v-model="model.email"
            placeholder="email"
            prefix-icon="el-icon-message"
          ></el-input>
        </el-form-item>
        <el-form-item prop="username">
          <el-input
            v-model="model.username"
            placeholder="username"
            prefix-icon="el-icon-user"
          ></el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input
            v-model="model.password"
            placeholder="Password"
            type="password"
            prefix-icon="el-icon-lock"
          ></el-input>
        </el-form-item>
        <el-form-item label="Gender" prop="gender">
          <el-radio-group v-model="model.gender">
            <el-radio label="Male"></el-radio>
            <el-radio label="Female"></el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item>
          <span class="el-form-item__label">Birthdate</span>
          <el-date-picker
            v-model="model.birthdate"
            type="date"
            placeholder="Pick a date"
            :default-value="new Date(2000, 1, 1)"
          >
          </el-date-picker>
        </el-form-item>
        <el-form-item>
          <el-button
            :loading="loading"
            class="signup-button"
            type="primary"
            native-type="submit"
            block
            >Sign up</el-button
          >
        </el-form-item>
      </el-form>
    </card>
    <hr />
    <card class="register-card box-shadow p18">
      <router-link to="/login">Have an account? Login</router-link>
    </card>
  </center>
</template>

<script>
import { mapActions } from 'vuex'

export default {
  name: 'signup',
  data() {
    return {
      model: {
        email: 'test1@gmail.com',
        username: 'Database Admin',
        password: '123asqbo',
        gender: '',
        birthdate: '',
      },
      loading: false,
      rules: {
        email: [
          { required: true, message: 'Email is required', trigger: 'blur' },
          { type: 'email', message: 'Email address invalid!', trigger: 'blur' },
        ],
        username: [
          { required: true, message: 'Username is required', trigger: 'blur' },
          {
            min: 8,
            message: 'Username must be at least 8 characters',
            trigger: 'blur',
          },
        ],
        password: [
          { required: true, message: 'Password is required', trigger: 'blur' },
          {
            min: 8,
            message: 'Password must be at least 8 characters',
            trigger: 'blur',
          },
        ],
      },
    }
  },
  methods: {
    ...mapActions(['signup']),
    onSingup() {
      let valid = this.$refs.form.validate()
      if (!valid) {
        return
      }
      this.loading = true
      this.signup(this.model).then(() => {
        this.loading = false
        this.$router.push({ name: 'Login' })
      })
    },
  },
}
</script>

<style scoped>
.signup .card {
  text-align: center;
  width: 300px;
  background: var(--hl);
}
.signup-button {
  width: 100%;
  margin-top: 30px;
}
h2 {
  letter-spacing: 1px;
  padding-bottom: 20px;
}
hr {
  margin-top: 20px;
  margin-bottom: 20px;
}
</style>
