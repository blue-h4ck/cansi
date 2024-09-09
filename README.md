# cansi

SuperT returns a string with applied color styles using ANSI escape sequences.

Parameters:
	- x: The string to be styled with colors.
- option: Integer that determines which type of color will be applied:
  - 0: Applies color to the text (first RGB value in 'rgb').
  - 1: Applies color to the background (first RGB value in 'rgb').
  - 2: Applies color to the text (first RGB value in 'rgb') and to the background (second RGB value in 'rgb').
- rgb: Slice of arrays of three integers representing the RGB values for the text and/or background color.
Returns:
- A new string with the applied text and/or background color, followed by a reset code to revert the colors.
Usage examples:
- SuperT("hello", 0, [3]int{255, 0, 0})   // Returns "hello" with red text color.
- SuperT("hello", 1, [3]int{0, 255, 0})   // Returns "hello" with green background color.
- SuperT("hello", 2, [3]int{255, 0, 0}, [3]int{0, 0, 255})   // Returns "hello" with red text color and blue background color.



SuperP applies color styles to a string using ANSI escape sequences.
Parameters:
- x: Pointer to the string to be modified with colors.
- option: Integer that determines which type of color will be applied:
  - 0: Applies color to the text (first RGB value in 'rgb').
  - 1: Applies color to the background (first RGB value in 'rgb').
  - 2: Applies color to the text (first RGB value in 'rgb') and to the background (second RGB value in 'rgb').
- rgb: Slice of arrays of three integers representing the RGB values for the text and/or background color.
Usage examples:
- SuperP(&text, 0, [3]int{255, 0, 0})   // Applies red color to the text.
- SuperP(&text, 1, [3]int{0, 255, 0})   // Applies green color to the background.
- SuperP(&text, 2, [3]int{255, 0, 0}, [3]int{0, 0, 255})   // Applies red color to the text and blue to the background.
The function also adds a reset code at the end to revert the colors after the text.
