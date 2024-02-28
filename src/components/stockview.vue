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
        :items="newitemadd"
        :search="search" 
        ></v-data-table>
    </v-card>
    {{ newarr }}
    <!-- <v-btn @click="click">click</v-btn> -->
    </div>
</template>
<script>
  export default {
    name:'',
    components:{
    },
    data () {
      return {
        search:'',
        
        medicine:this.$store.state.medicinemaster,
        stock:this.$store.state.stock,
        // mergedArray:this.$store.state.stock.concat(this.$store.state.medicinemaster),
        headers: [ 
          // {text:'Medicine Name'}, {text:'Brand Name'},{text:'Qty'},{text:'Unit Price'}
          {
          text: "Medicine Name",
          align: "start",
          sortable: false,
          value: "medicinename",
        },
        { text: "Brand Name", value: "brandname" },
        { text: "Qty", value: "quantity" },
        { text: "Unit Price", value: "amount" }, 
        ],
        newarr:[]
        
      }
    },
    props:{
      // sendvalue:String
    },
    methods:{
    },
    watch:{

    newarr:{
      handler(){
        if (this.newarr.length>=0){
          console.log('kkkk')        
        }
      },immediate:true
    }
    },
    computed:{
      newitemadd(){
        let newarr=[]
        for(var i in this.stock){
            for(var j in this.medicine){
                if(this.stock[i].medicinename === this.medicine[j].medicinename){
                    let mergedObject = {
                        medicinename: this.stock[i].medicinename,
                        quantity: this.stock[i].quantity, 
                        amount: this.stock[i].amount, 
                        brandname: this.medicine[j].brandname
                    }
                    newarr.push(mergedObject)
                }
            }
          }  
        return newarr
      }
    }

  }
</script>

