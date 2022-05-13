import { useState } from "react"

const Basket = () => {
    var [bascket, setBasket] = useState(sessionStorage.getItem("basket") == null ? [] : JSON.parse(sessionStorage.getItem("basket")))

    const buyForProducts = (event) => {

    }

    return (
        <div className="App">
            <table>
                {
                    bascket.map((element, i) => {
                        return <tr><td>element.product_name</td><td>element.quantity</td></tr>
                    })
                }
            </table>
            <button onClick={buyForProducts}>Pay</button>
        </div>
    )
}

export default Basket