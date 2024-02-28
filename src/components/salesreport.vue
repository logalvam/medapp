<template>
    <div>
        <v-card width="500" height="600" class="green">
            <!-- {{ stock }} -->
            <!-- {{ medicine }} -->
            {{ newarr }}
        </v-card>
        <v-btn @click="show">click</v-btn>
    </div>
</template>
<script>
export default{
    name:'',
    data(){
        return{
            stock:this.$store.state.stock,
            medicine:this.$store.state.medicinemaster,
            newarr:[]
        }
    },
    methods:{
        show(){
            for(var i in this.stock){
                for(var j in this.medicine){
                    if(this.stock[i].medicinename === this.medicine[j].medicinename){
                        let mergedObject = {
                            medicinename: this.stock[i].medicinename,
                            quantity: this.stock[i].quantity, // Assuming this property exists in stock
                            amount: this.stock[i].amount, // Assuming this property exists in stock
                            brandname: this.medicine[j].brandname
                        }
                        this.newarr.push(mergedObject)
                    }
                }
            }
            // this.newarr=Array.prototype.push.apply(this.stock,this.medicine);
        }
    }
}
</script>