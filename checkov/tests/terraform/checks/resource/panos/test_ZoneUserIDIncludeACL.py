import unittest
import os

from checkov.terraform.checks.resource.panos.ZoneUserIDIncludeACL import check
from checkov.runner_filter import RunnerFilter
from checkov.terraform.runner import Runner


class TestZoneUserIDIncludeACL(unittest.TestCase):

    def test(self):
        runner = Runner()
        current_dir = os.path.dirname(os.path.realpath(__file__))

        test_files_dir = current_dir + "/example_ZoneUserIDIncludeACL"
        report = runner.run(root_folder=test_files_dir, runner_filter=RunnerFilter(checks=[check.id]))
        summary = report.get_summary()

        passing_resources = {
            'panos_zone.pass1',
            'panos_panorama_zone.pass2',
            'panos_zone.pass3',
            'panos_panorama_zone.pass4',
            'panos_zone.pass5',
            'panos_panorama_zone.pass6',
        }
        failing_resources = {
            'panos_zone.fail1',
            'panos_panorama_zone.fail2',
            'panos_zone.fail3',
            'panos_panorama_zone.fail4',
            'panos_zone.fail5',
            'panos_panorama_zone.fail6',
        }

        passed_check_resources = set([c.resource for c in report.passed_checks])
        failed_check_resources = set([c.resource for c in report.failed_checks])

        self.assertEqual(summary['passed'], 6)
        self.assertEqual(summary['failed'], 6)
        self.assertEqual(summary['skipped'], 0)
        self.assertEqual(summary['parsing_errors'], 0)

        self.assertEqual(passing_resources, passed_check_resources)
        self.assertEqual(failing_resources, failed_check_resources)


if __name__ == '__main__':
    unittest.main()
