<template>
  <div>
    <v-card>
      <v-card-title>
        <v-text-field
          v-model="search"
          append-icon="mdi-magnify"
          label="Search"
          single-line
          hide-details
        ></v-text-field>
      </v-card-title>

      <v-data-table
        :headers="headers"
        :items="medicine"
        :search="search"
      ></v-data-table>
    </v-card>
    <div>{{ stock }}</div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      medicine: this.$store.state.medicinemaster,
      stock: this.$store.state.stock,
      headers: [
        // {text:'Medicine Name'}, {text:'Brand Name'},{text:'Qty'},{text:'Unit Price'}
        {
          text: "Medicine Name",
          align: "start",
          sortable: false,
          value: "name",
        },
        { text: "Brand Name", value: "calories" },
        { text: "Qty", value: "fat" },
        { text: "Unit Price", value: "carbs" },
      ],
    };
  },
};
</script>
