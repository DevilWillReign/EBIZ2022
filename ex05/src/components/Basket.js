import { useState } from "react"

const Basket = () => {
    var [bascket, setBasket] = useState([])

    return (
        <div className="App">
            <table>
                {
                    bascket.map((element, i) => {
                        return <tr><td>element.product_name</td><td>element.quantity</td></tr>
                    })
                }
            </table>
        </div>
    )
}

export default Basket