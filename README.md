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

## using intlij plugin

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