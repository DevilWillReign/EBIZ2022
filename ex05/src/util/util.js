
const myRange = (lenght, start = 0) => {
    return [...Array(lenght).keys()].map(i => i + start)
}

export default (
    myRange
)