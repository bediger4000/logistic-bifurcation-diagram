# Logistic Bifurcation Diagram

Generate an image of the bifurcation diagram of the logistic equation

```none
x1 = r * x0 * (1 - x0)
```

This is about as cheap and dirty as I could get.
I use [gnuplot](http://www.gnuplot.info/) to draw the image,
I didn't even bother to have the Go image package involved.
My program outputs text representations of floating point numbers.

## Algorithm

0. For values of r from 2.8 to 4.0 (step 0.0002):
1. Set x = 0.5
2. Run 100 iterations of x = r * x * (1.0 - x)
3. Run 100 more iterations of x = r * x * (1.0 - x),
this time keeping all (floating point!) values of x
4. Print out r and the unique values of x from step 3

Things that could change the diagram:

* Initial value of x
* Step size of r
* Number of "settling" iterations
* Number of iterations done to find unique values of x

My diagram has bulkier bifuration points than the Wikipedia diagrams.
I could not make it as graceful as other people got.

## References

* [Wikipedia](https://en.wikipedia.org/wiki/Bifurcation_diagram)
* [Python code](https://geoffboeing.com/2015/03/chaos-theory-logistic-map/)
* [Logistic map](https://en.wikipedia.org/wiki/Logistic_map)
