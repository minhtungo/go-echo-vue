<template>
  <div id="app">
    <div class="row">
      <div class="col-md-6 offset-md-3 py-5">
        <h1>Hello Taubyte</h1>
        <b-form @submit="postreq">
          <b-form-group label="CID:" label-for="input-1">
            <b-form-input
              id="input-1"
              v-model="form.cid"
              placeholder="Enter a CID"
              required
            ></b-form-input>
          </b-form-group>
          <br />
          <b-button type="submit" variant="primary">Submit</b-button>
        </b-form>
      </div>
    </div>
  </div>
</template>

<script>
// import axios from 'axios'; 

export default {
  data() {
    return {
      form: {
        cid: ""
      }
    }
  },
  methods: {
    // postreq() {
    //   var data = {"cid": string(this.cid)};
      
    // }
  }
};
</script>

<style></style>
