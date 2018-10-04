<template>
    <div>
        <listing :tableName="$route.params.child" :parent="$route.params.fid"></listing>             
    </div>
</template>