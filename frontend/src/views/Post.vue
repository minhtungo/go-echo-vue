<template>
  <div>
    <div class="row">
      <div class="col-md-6 offset-md-3 py-5">
        <h1>Hello Taubyte</h1>
        <b-form @submit="postreq">
          <b-form-group label="Content:" label-for="input-2">
            <b-form-input
              id="input-2"
              v-model="form.content"
              placeholder="Enter content"
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
      <div class="card-body">{{returnedCID}}</div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      form: {
        content: ""
      }
    };
  },
  methods: {
    postreq() {
      var content = {"content": this.form.content};
      axios.post("http://localhost:8000/add", content)
      .then(response => {this.returnedCID = response.data.cid})
      .catch(error => {
          this.errorMessage = error.message;
          console.error("There was an error!", error);
        });
      // }
    }
  }
};
</script>

<style></style>
