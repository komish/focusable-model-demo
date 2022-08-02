# Focusable Model Demo

This is a proof-of-concept testing the idea of a `Focusable` interface,
which allows you to pass "focus" to independently developed components, or `tea.Model`s.

Here, the `FocusableBoxArray` model implements this `Focusable` interface. Then
a `FocusController` model is built with multiple Focusable items, and
orchestrates passing focus across the various components.

This demo has KeyLeft and KeyRight messages being handled by the focus
controller. Other messages are passed directly to the underlying Focusable
component.

There are components that are not focusable but may need to be updated based on
actions taken on focusable components. This PoC stores those values in a Viper config
which is then retrieved by those components.
