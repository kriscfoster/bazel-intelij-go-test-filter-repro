# bazel-go-test-filter-repro

Reproducing some un-expected behaviour from Intelij bazel plugin.


## command line

```
➜ bazel test //... --test_filter=TestMultipleCases/TestCase1 
...
==================== Test output for //:example_test:
=== RUN   TestMultipleCases
2022/06/16 07:15:01 running test: TestCase1
=== RUN   TestMultipleCases/TestCase1
--- PASS: TestMultipleCases (0.00s)
    --- PASS: TestMultipleCases/TestCase1 (0.00s)
PASS
...
```

## using intelij plugin

<img width="842" alt="Screenshot 2022-06-16 at 08 21 01" src="https://user-images.githubusercontent.com/17026751/174014898-0832bfed-fa38-4c21-82fe-e7ef2887118c.png">

The correct `--test_filter` isn't applied & more than the desired tests run:

```
/Users/kfoster/bin/bazel test --tool_tag=ijwb:IDEA:ultimate ... --test_filter=^TestMultipleCases$ -- //:example_test
...
==================== Test output for //:example_test:
=== RUN   TestMultipleCases
2022/06/16 07:18:22 running test: TestCase1
2022/06/16 07:18:22 running test: TestCase2
=== RUN   TestMultipleCases/TestCase1
=== RUN   TestMultipleCases/TestCase2
--- PASS: TestMultipleCases (0.00s)
    --- PASS: TestMultipleCases/TestCase1 (0.00s)
    --- PASS: TestMultipleCases/TestCase2 (0.00s)
PASS
================================================================================
Target //:example_test up-to-date:
```
