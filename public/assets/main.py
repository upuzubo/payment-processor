# main.py

from payment_processor import PaymentProcessor

if __name__ == "__main__":
    payment_processor = PaymentProcessor()
    payment_processor.initiate_payment("customer_id", 10.99)
    payment_processor.capture_payment("payment_id")
    payment_processor.refund_payment("payment_id", 5.00)
    payment_processor.cancel_payment("payment_id")