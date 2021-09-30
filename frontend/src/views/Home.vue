<template>
  <div id="app">
    <div class="row">
      <div class="col-md-6 offset-md-3 py-5">
        <h1>Get Content from IPFS</h1>
        <b-form @submit.prevent="getFormValues">
          <b-form-group label="CID:" label-for="input-1">
            <b-form-input
              id="input-1"
              v-model="cid"
              placeholder="Enter a CID:"
              required
            ></b-form-input>
          </b-form-group>
          <br />
          <b-button type="submit" variant="primary">Submit</b-button>
        </b-form>
      </div>
    </div>
    <div class="card text-center m-3">
      <h5 class="card-header">Content</h5>
      <div class="card-body">{{ content }}</div>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      content: null,
    };
  },
  methods: {
    getFormValues() {
      axios
        .get("http://localhost:8000/get?cid=" + this.cid)
        .then((response) => (this.content = response.data.data))
        .catch((error) => {
          this.errorMessage = error.message;
          console.error("There was an error!", error);
        });
    },
  },
};
</script>

<style></style>
