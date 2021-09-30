<template>
  <div id="app">
    <div class="row">
      <div class="col-md-6 offset-md-3 py-5">
        <h1>Post Content to IPFS</h1>
        <b-form @submit.prevent="getFormValues">
          <b-form-group label="CID:" label-for="input-1">
            <b-form-input
              id="input-1"
              v-model="content"
              placeholder="Enter content:"
              required
            ></b-form-input>
          </b-form-group>
          <br />
          <b-button type="submit" variant="primary">Submit</b-button>
        </b-form>
      </div>
    </div>
    <div class="card text-center m-3">
      <h5 class="card-header">CID</h5>
      <div class="card-body">{{ cid }}</div>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      cid: null,
    };
  },
  methods: {
    getFormValues() {
      axios
        .post("http://localhost:8000/add", this.content)
        .then((response) => (this.cid = response.data.cid))
        .catch((error) => {
          this.errorMessage = error.message;
          console.error("There was an error!", error);
        });
    },
  },
};
</script>

<style></style>
