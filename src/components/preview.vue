<template>
    <div>
        <v-sheet width="">

            <v-container class="green" >
                <v-row>
                    <v-col cols="12" md="12">
                        <h1 class="text-center">Medical Shop Name</h1>
                    </v-col>
                    <v-divider></v-divider>
                    <v-col cols="12" md="7" class="white d-flex  justify-space-between">
                        <h1>Medicine name</h1>
                        <h1>Qty</h1>
                        <h1>Amount</h1>
                    </v-col>
                    <!-- <v-col v-for="(item,value) in items" :key="item"> -->

                    <!-- </v-col> -->
                </v-row>
            </v-container>
        </v-sheet>
            
    </div>
</template>
<script>
export default{
    name:'',
    data(){
        return{

        }
    }
}
</script>