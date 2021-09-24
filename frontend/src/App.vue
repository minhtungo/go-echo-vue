<template>
  <div id="app">
    <div id="nav">
      <b-navbar type="dark" variant="dark">
        <b-navbar-nav>
          <b-nav-item><router-link to="/">Home</router-link> </b-nav-item>
          <b-nav-item><router-link to="/post">Post</router-link></b-nav-item>
        </b-navbar-nav>
      </b-navbar>
    </div>
    <router-view />
  </div>
</template>

<style></style>
