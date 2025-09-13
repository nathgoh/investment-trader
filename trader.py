import streamlit as st
import requests

st.title("AI Investment Trader")

def get_health():
    """
    Calls the Go backend to get the health status.
    """

    try:
        response = requests.get("http://localhost:8080/api/v1/health")
        response.raise_for_status()  # Raise an exception for bad status codes
        return response.json()
    except requests.exceptions.RequestException as e:
        st.error(f"Error calling the backend: {e}")
        return None

def get_stock_quote(symbol):
    """
    Calls the Go backend to get the stock quote for a given symbol.
    """

    try:
        response = requests.get(f"http://localhost:8080/api/marketdata/quotes/{symbol}")
        response.raise_for_status()  # Raise an exception for bad status codes
        return response.json()
    except requests.exceptions.RequestException as e:
        st.error(f"Error calling the backend: {e}")
        return None

def stock_prediction():
    st.subheader("Stock Quote")
    st.write("This section will display the latest stock quote.")

    stock_symbol = st.text_input("Enter Stock Symbol (e.g. AAPL)")

    if st.button("Get Quote"):
        if stock_symbol:
            st.write(f"Fetching quote for {stock_symbol}...")
            quote = get_stock_quote(stock_symbol)
            if quote:
                st.json(quote)
        else:
            st.warning("Please enter a stock symbol.")

    if st.button("Get Health"):
        health = get_health()
        if health:
            st.json(health)

stock_prediction()
