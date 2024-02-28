<template>
  <div>
    <v-app-bar color="blue-grey lighten-1" elevation="4">
      <v-container fluid class="white">
        <v-row class="d-flex flex-row pa-3 justify-end">
          <!-- <v-col class=""> -->
          <v-btn class="ml-3">Dashboard</v-btn>
          <v-btn v-show="addUser" class="ml-3" @click="adduser">Adduser</v-btn>
          <v-btn v-show="stockView" @click="stockview" class="ml-3"
            >StockView</v-btn
          >
          <v-btn v-show="stockEntry" @click="stockentry" class="ml-3">stock Entry</v-btn>
          <v-btn v-show="billeEntry" @click="bill" class="ml-3"
            >Bill Entry</v-btn
          >
          <v-btn v-show="SaleReport" class="ml-3">Sales Report</v-btn>
          <v-btn v-show="loginHistory" class="ml-3">login history</v-btn>
          <v-btn class="ml-3" @click="logout">login out</v-btn>

          <!-- </v-col> -->
        </v-row>
      </v-container>
    </v-app-bar>
    <!-- Adduser -->
    <div class="blue">
      <!-- <h1>{{ date }} 666</h1> -->
    </div>
  </div>
</template>
<script>
export default {
  name: "",
  data() {
    return {
      addUser: true,
      stockEntry: true,
      stockView: true,
      billeEntry: true,
      SaleReport: true,
      loginHistory: true,
    };
  },
  props: {
    role: String,
    date: Date,
    //   billeEntry:Boolean,
    //   stockEntry:Boolean,
    //   stockView:Boolean,
    //   SaleReport:Boolean,
    //   addUser:Boolean,
    //   loginHistory:Boolean
  },
  methods: {
    logout() {
      this.$router.push("/");
    },
    adduser() {
      this.$router.push("/adduser");
    },
    bill() {
      this.$router.push("/bill");
    },
    stockview() {
      this.$router.push("/stockView1");
    },
    stockentry() {
        this.$router.push("/stockentry")
    }
  },
  watch: {
    role() {
      if (this.role === "Biller") {
        alert("biller");
      }
    },
  },
};
</script>
