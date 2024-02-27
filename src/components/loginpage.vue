<template>
     <div >
            <v-card flat rounded="false" height="100vh">

                <v-row justify="center" align='center'>
                    <v-card width="400" height="300" color="">
                        <v-container fluid>
                            <v-row>
                            <v-col cols="12" md-12>
                                <v-text-field v-model="userid" label="userid">
                                </v-text-field>
                            </v-col>
                            <v-col cols="12" md-12>
                                <v-text-field v-model="password" label="password">
                                </v-text-field>
                            </v-col>
                            <v-col justify="center">
                                <v-btn md="5" color="primary">
                                    login
                                </v-btn>
                            </v-col>
                        </v-row>
                        </v-container>
                    </v-card>
                </v-row>
            </v-card >
        </div>
</template>
<script>
export default{
    name:'loginpage',
    date(){
        return{
            userid:'',
            password:''
        }
    }
}
</script>