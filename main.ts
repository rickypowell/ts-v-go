
function getMyData() {
    return Promise.resolve(`{ "name": "ricky", "age": 99 }`);
}

type MyData = {
    age: number;
};

async function main() {
    try {

        const payload = await getMyData();
        console.log('before', payload);
        const result: MyData = JSON.parse(payload);

        const { age } = result;

        console.log(`${age} ${typeof age}`);

        const transformedData = JSON.stringify({
            ...result,
            age: 4
        });

        console.log('after', transformedData);
    } catch (error) {
        // error occurred
    }
}

main();
